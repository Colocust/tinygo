package tinygo

import (
	"log"
	"net/http"
	"sync"
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
)

func Default() *Engine {
	e := New()
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

	engine.router[http.MethodGet] = make(map[string]HandlersChain)
	engine.router[http.MethodPost] = make(map[string]HandlersChain)

	engine.RouterGroup.engine = engine
	return engine
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	ctx := new(Context)
	ctx.Request, ctx.Response = req, writer
	ctx.reset()
	e.handleHTTPRequest(ctx)
}

func (e *Engine) handleHTTPRequest(ctx *Context) {
	method, uri := ctx.Request.Method, ctx.Request.URL.Path
	handlers := e.findRouteByUri(method, uri)
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

func (e *Engine) findRouteByUri(method string, uri string) HandlersChain {
	log.Println(uri, e.router)
	return e.router[method][uri]
}

func (e *Engine) addRoute(method string, uri string, handlers HandlersChain) {
	e.router[method][uri] = handlers
}

func serveError(ctx *Context, status int, defaultMessage []byte) {

}
