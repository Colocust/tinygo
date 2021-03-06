package tinygo

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Engine struct {
	router map[string]map[string]HandlersChain
	RouterGroup
	pool       sync.Pool
	allNoRoute HandlersChain
}

type HandlerFunc func(c *Context)
type HandlersChain []HandlerFunc

var (
	default404Body = []byte("404 page not found")
	default504Body = "504 Gateway Time-out"
)

func Default() *Engine {
	e := New()
	e.Use(HandlerFuncWithTimeout(time.Millisecond * 500))
	return e
}

func New() *Engine {
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
		},
		router: make(map[string]map[string]HandlersChain),
	}
	engine.routerInit()
	engine.RouterGroup.engine = engine

	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
	return engine
}

func (e *Engine) allocateContext() *Context {
	return &Context{
		engine: e,
		Writer: &responseWriter{},
	}
}

func (e *Engine) routerInit() {
	e.router[http.MethodGet] = make(map[string]HandlersChain)
	e.router[http.MethodPost] = make(map[string]HandlersChain)
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	ctx := e.pool.Get().(*Context)

	ctx.Writer.Reset(writer)
	ctx.Request = req
	ctx.reset()

	e.handleHTTPRequest(ctx)
}

func (e *Engine) handleHTTPRequest(ctx *Context) {
	method, uri := ctx.Request.Method, ctx.Request.URL.Path
	handlers := e.findHandlersByUri(method, uri)
	if len(handlers) > 0 {
		ctx.Handlers = handlers
		ctx.Next()
		return
	}

	ctx.Handlers = e.allNoRoute
	serveError(ctx, http.StatusNotFound, default404Body)
}

func (e *Engine) Run(addr string) (err error) {
	err = http.ListenAndServe(addr, e)
	return
}

func (e *Engine) findHandlersByUri(method string, uri string) HandlersChain {
	return e.router[method][uri]
}

func (e *Engine) addRoute(method string, uri string, handlers HandlersChain) {
	e.router[method][uri] = handlers
}

var mimePlain = []string{"text/plain"}

func serveError(ctx *Context, status int, defaultMessage []byte) {
	if ctx.Writer.Written() {
		return
	}

	ctx.Writer.Header()["Content-Type"] = mimePlain

	ctx.Writer.WriteHeader(status)
	_, _ = ctx.Writer.Write(defaultMessage)
}

func HandlerFuncWithTimeout(t time.Duration) HandlerFunc {
	return func(c *Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)

		// ?????????????????? ???????????????????????????
		finish := make(chan struct{}, 1)
		go func() {
			c.Next()
			finish <- struct{}{}
		}()

		select {
		case <-ctx.Done():
			c.Abort()
			c.Json(http.StatusGatewayTimeout, default504Body)
			c.SetTimeout()
		case <-finish:
		}
	}
}
