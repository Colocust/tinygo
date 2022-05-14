package app

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
}

// 为当前路由组添加中间件
func (group *RouterGroup) Use(middlewares ...HandlerFunc) Routes {
	group.Handlers = append(group.Handlers, middlewares...)
	return group.returnObj()
}

func (group *RouterGroup) Group(path string, middlewares ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		Handlers: group.combineHandlers(middlewares),
		basePath: joinPaths(group.basePath, path),
		engine:   group.engine,
	}
}

func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	finalSize := len(group.Handlers) + len(handlers)
	mergedHandlers := make(HandlersChain, finalSize)
	copy(mergedHandlers, group.Handlers)
	copy(mergedHandlers[len(group.Handlers):], handlers)
	return mergedHandlers
}

func (group *RouterGroup) returnObj() Routes {
	if group.isRoot {
		return group.engine
	}
	return group
}
