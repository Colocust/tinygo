package app

import "sync"

type Engine struct {
	RouterGroup               // 初始路由组
	noRoute     HandlersChain // 404Handler
	noMethod    HandlersChain // 405Handler
	pool        sync.Pool     // 资源池
	trees       methodsTree   // 路由树
}

func New() *Engine {
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
			isRoot:   true,
		},
		trees: make(methodsTree, 0, 9),    // GET POST DELETE PUT HEAD PATCH OPTIONS CONNECT TRACE
	}

	engine.RouterGroup.engine = engine

	return engine
}

func Default() *Engine {
	engine := New()

	return engine
}
