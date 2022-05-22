package tinygo

import (
	"context"
	"net/http"
	"sync"
)

type Engine struct {
	context.Context
	router map[string]map[string]HandlersChain
	RouterGroup
	pool sync.Pool
}

type HandlerFunc func(c *Context) error
type HandlersChain []HandlerFunc

func (e *Engine) ServeHTTP(writer http.ResponseWriter, req *http.Request) {

}

func (e *Engine) addRoute(method string, absolutePath string, handlers HandlersChain) {

}

func (e *Engine) findRouteByRequest(uri string) HandlersChain {

}

func (e *Engine) Get(path string, handlers ...HandlerFunc) IRoutes {
	return e.RouterGroup.handle(http.MethodGet, path, handlers...)
}

func (e *Engine) Use(handlers ...HandlerFunc) IRoutes {
	e.RouterGroup.Use(handlers...)
	return e
}
