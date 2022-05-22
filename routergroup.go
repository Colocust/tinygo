package tinygo

import "net/http"

type IRoutes interface {
	Group(string, ...HandlerFunc) IRoutes
	Use(...HandlerFunc) IRoutes
	Post(string, ...HandlerFunc) IRoutes
	Get(string, ...HandlerFunc) IRoutes
}

type RouterGroup struct {
	engine   *Engine
	basePath string
	Handlers HandlersChain
}

func (rg *RouterGroup) Group(path string, handlers ...HandlerFunc) IRoutes {
	return &RouterGroup{
		engine:   rg.engine,
		basePath: rg.calculateAbsolutePath(path),
		Handlers: rg.combineHandlers(handlers...),
	}
}

func (rg *RouterGroup) Use(handlers ...HandlerFunc) IRoutes {
	rg.Handlers = append(rg.Handlers, handlers...)
	return rg
}

func (rg *RouterGroup) handle(method string, path string, handlers ...HandlerFunc) IRoutes {
	absolutePath, finalHandlers := rg.calculateAbsolutePath(path), rg.combineHandlers(handlers...)
	rg.engine.addRoute(method, absolutePath, finalHandlers)
	return rg
}

func (rg *RouterGroup) Get(path string, handlers ...HandlerFunc) IRoutes {
	return rg.handle(http.MethodGet, path, handlers...)
}

func (rg *RouterGroup) Post(path string, handlers ...HandlerFunc) IRoutes {
	return rg.handle(http.MethodPost, path, handlers...)
}

func (rg *RouterGroup) combineHandlers(handlers ...HandlerFunc) HandlersChain {
	finalSize := len(rg.Handlers) + len(handlers)
	finalHandlers := make([]HandlerFunc, finalSize)
	copy(finalHandlers, rg.Handlers)
	copy(finalHandlers[len(rg.Handlers):], handlers)
	return finalHandlers
}

func (rg *RouterGroup) calculateAbsolutePath(path string) string {
	return joinPaths(rg.basePath, path)
}
