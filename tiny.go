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
	engine.routerReset()
	engine.RouterGroup.engine = engine

	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
	return engine
}

func (e *Engine) allocateContext() *Context {
	return &Context{
		engine: e,
	}
}

func (e *Engine) routerReset() {
	e.router[http.MethodGet] = make(map[string]HandlersChain)
	e.router[http.MethodPost] = make(map[string]HandlersChain)
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	ctx := e.pool.Get().(*Context)
	ctx.Writer = &responseWriter{
		ResponseWriter: writer,
	}
	ctx.Writer.Reset()
	ctx.Request = req

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

func serveError(ctx *Context, status int, defaultMessage []byte) {

}

func HandlerFuncWithTimeout(t time.Duration) HandlerFunc {
	return func(c *Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)

		// 一定要有容量 否则子协程无法退出
		finish := make(chan struct{}, 1)
		go func() {
			c.Next()
			finish <- struct{}{}
		}()

		select {
		case <-ctx.Done():
			c.Abort()
			c.Json(http.StatusGatewayTimeout, default504Body)
		case <-finish:
		}
	}
}
