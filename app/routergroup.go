package app

import "net/http"

type RouterGroup struct {
	Handlers HandlersChain // 路由组的全部中间件
	basePath string        // 路由组path
	engine   *Engine       // 所属engine
	isRoot   bool          // 是否为根路由组
}

type Router interface {
	Group(string, ...HandlerFunc) *RouterGroup
	Routes
}

// 定义了路由所有的接口
type Routes interface {
	Use(...HandlerFunc) Routes

	Get(string, ...HandlerFunc) Routes
	Head(string, ...HandlerFunc) Routes
	Post(string, ...HandlerFunc) Routes
	Put(string, ...HandlerFunc) Routes
	Patch(string, ...HandlerFunc) Routes
	Delete(string, ...HandlerFunc) Routes
	Connect(string, ...HandlerFunc) Routes
	Options(string, ...HandlerFunc) Routes
	Trace(string, ...HandlerFunc) Routes
}

// 为当前路由组添加中间件
func (group *RouterGroup) Use(handlers ...HandlerFunc) Routes {
	group.Handlers = append(group.Handlers, handlers...)
	return group.returnObj()
}

func (group *RouterGroup) Group(path string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		Handlers: group.combineHandlers(handlers),
		basePath: group.calculateAbsolutePath(path),
		engine:   group.engine,
	}
}

func (group *RouterGroup) BasePath() string {
	return group.basePath
}

func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) Routes {
	absolutePath := group.calculateAbsolutePath(relativePath)
	handlers = group.combineHandlers(handlers)
	group.engine.addRoute(httpMethod, absolutePath, handlers)
	return group.returnObj()
}

func (group *RouterGroup) Get(path string, handlers ...HandlerFunc) Routes {
	return group.handle(http.MethodGet, path, handlers)
}

func (group *RouterGroup) Post(path string, handlers ...HandlerFunc) Routes {
	return group.handle(http.MethodPost, path, handlers)
}

func (group *RouterGroup) Head(path string, handlers ...HandlerFunc) Routes {
	return group.handle(http.MethodHead, path, handlers)
}

func (group *RouterGroup) Put(path string, handlers ...HandlerFunc) Routes {
	return group.handle(http.MethodPut, path, handlers)
}

func (group *RouterGroup) Patch(path string, handlers ...HandlerFunc) Routes {
	return group.handle(http.MethodPatch, path, handlers)
}

func (group *RouterGroup) Delete(path string, handlers ...HandlerFunc) Routes {
	return group.handle(http.MethodDelete, path, handlers)
}

func (group *RouterGroup) Connect(path string, handlers ...HandlerFunc) Routes {
	return group.handle(http.MethodConnect, path, handlers)
}

func (group *RouterGroup) Options(path string, handlers ...HandlerFunc) Routes {
	return group.handle(http.MethodOptions, path, handlers)
}

func (group *RouterGroup) Trace(path string, handlers ...HandlerFunc) Routes {
	return group.handle(http.MethodTrace, path, handlers)
}

func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	finalSize := len(group.Handlers) + len(handlers)
	mergedHandlers := make(HandlersChain, finalSize)
	copy(mergedHandlers, group.Handlers)
	copy(mergedHandlers[len(group.Handlers):], handlers)
	return mergedHandlers
}

func (group *RouterGroup) calculateAbsolutePath(relativePath string) string {
	return joinPaths(group.basePath, relativePath)
}

func (group *RouterGroup) returnObj() Routes {
	if group.isRoot {
		return group.engine
	}
	return group
}
