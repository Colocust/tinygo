package tinygo

import (
	"net/http"
	"sync"
)

type Engine struct {
	RouterGroup               // 初始路由组
	allNoRoute  HandlersChain // 404 handlers
	allNoMethod HandlersChain // 405 handlers
	noRoute     HandlersChain // 设置NoRoute时的404 handlers
	noMethod    HandlersChain // 设置NoMethod时的404 handlers
	pool        sync.Pool     // 资源池
	trees       methodsTree   // 路由树
}

type HandlerFunc func(ctx *Context)

type HandlersChain []HandlerFunc

var (
	default404Body = []byte("404 page not found")
)

func New() *Engine {
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
			isRoot:   true,
		},
		trees: make(methodsTree, 0, 9), // GET POST DELETE PUT HEAD PATCH OPTIONS CONNECT TRACE
	}

	engine.RouterGroup.engine = engine
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}

	return engine
}

func Default() *Engine {
	engine := New()
	engine.Use()
	return engine
}

func (engine *Engine) Use(handlers ...HandlerFunc) Routes {
	engine.RouterGroup.Use(handlers...)
	engine.rebuild404Handlers()
	engine.rebuild405Handlers()
	return engine
}

func (engine *Engine) Run(addr string) (err error) {
	err = http.ListenAndServe(addr, engine)
	return
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := engine.pool.Get().(*Context)
	ctx.Writer.reset(w)
	ctx.Reset()
	ctx.Request = req

	engine.handleHTTPRequest(ctx)
	engine.pool.Put(ctx)
}

func (engine *Engine) handleHTTPRequest(ctx *Context) {
	method, _, trees := ctx.Request.Method, ctx.Request.URL.Path, engine.trees
	for i, tl := 0, len(trees); i < tl; i++ {
		if method != trees[i].method {
			continue
		}
		//root := trees[i].root
		//value := root.getValue(path)
	}

	ctx.handlers = engine.allNoRoute
	serveError(ctx, http.StatusNotFound, default404Body)
}

func serveError(ctx *Context, status int, defaultMessage []byte) {
	ctx.Writer.status = status
	ctx.Next()
	if ctx.Writer.Written() {
		return
	}
	ctx.Writer.Write(defaultMessage)
}

func (engine *Engine) NoRoute(handlers ...HandlerFunc) {
	engine.noRoute = handlers
	engine.rebuild404Handlers()
}

func (engine *Engine) NoMethod(handlers ...HandlerFunc) {
	engine.noMethod = handlers
}

// 申请context资源
func (engine *Engine) allocateContext() *Context {
	return &Context{engine: engine, Writer: &ResponseWriter{}}
}

func (engine *Engine) rebuild404Handlers() {
	engine.allNoRoute = engine.combineHandlers(engine.noRoute)
}

func (engine *Engine) rebuild405Handlers() {
	engine.allNoMethod = engine.combineHandlers(engine.noMethod)
}

func (engine *Engine) addRoute(httpMethod, path string, handlers HandlersChain) {

}
