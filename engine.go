package tinygo

import "sync"

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

func (engine *Engine) NoRoute(handlers ...HandlerFunc) {
	engine.noRoute = handlers
	engine.rebuild404Handlers()
}

func (engine *Engine) NoMethod(handlers ...HandlerFunc) {
	engine.noMethod = handlers
}

func (engine *Engine) rebuild404Handlers() {
	engine.allNoRoute = engine.combineHandlers(engine.noRoute)
}

func (engine *Engine) rebuild405Handlers() {
	engine.allNoMethod = engine.combineHandlers(engine.noMethod)
}

func (engine *Engine) addRoute(httpMethod, path string, handlers HandlersChain) {

}
