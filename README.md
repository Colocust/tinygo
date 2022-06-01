# Gin框架学习

## 开篇
Golang想要实现一个简易的Web框架非常简单，只需要实现Http包中的Handler接口即可。
```golang
package framework

import "net/http"

type Framework struct {
	
}

func (f *Framework) Run () {
	address := "127.0.0.1:80"
	if err := http.ListenAndServe(address, f);err != nil {
		panic(err)
    }
}

func (f *Framework) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 获取req里的信息（uri、参数、http头部信息等）
	// 通过uri获取需要执行的函数并执行
	// 写入http状态码以及返回体
}
```
就像上面那样，在ServeHTTP函数承载自己框架的逻辑就可以了。不过在这之前，我们一般还需要注册路由以及中间件等。

## 路由组

### 结构
```go
package gin

type RouterGroup struct {
    Handlers HandlersChain   // 路由组绑定的中间件
    basePath string          // 路由组的路径
    engine   *Engine         // 路由组所属engine
    root     bool            // 是否为根路由
}
```

* Handlers为当前路由组绑定的所有中间件。当我们为当前路由组添加路由或者基于此创建一个新的路由组时，会将这些中间件添加到路由或新的路由组中。
* basePath与Handler类似，会基于basePath生成一个新的path。
* engine指针指向当前路由组的engine容器。当我们为路由组添加一个新的路由时，也会向该容器中的methodTrees中添加一个路由节点。
* root标记当前是否为根路由组。

### 使用

这里我们以创建一个新的路由组为例：
```go
package gin 

func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
    return &RouterGroup{
        Handlers: group.combineHandlers(handlers),
        basePath: group.calculateAbsolutePath(relativePath),
        engine:   group.engine,
    }
}

func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
    finalSize := len(group.Handlers) + len(handlers)
    assert1(finalSize < int(abortIndex), "too many handlers")
    mergedHandlers := make(HandlersChain, finalSize)
    copy(mergedHandlers, group.Handlers)
    copy(mergedHandlers[len(group.Handlers):], handlers)
    return mergedHandlers
}

func (group *RouterGroup) calculateAbsolutePath(relativePath string) string {
    return joinPaths(group.basePath, relativePath)
}
```
Group函数就是在当前路由组的基础上创建一个新的路由组了，创建时会将当前路由组的所有中间件与新路由组的中间件合并在一起并绑定在新的路由组结构中，同理，也会基于basePath拼接一个新的path。
```go
g1 := gin.RouterGroup{
    Handlers: gin.HandlersChain{
        gin.Logger(),
    },
    basePath: "/",
}
g2 := g1.Group("user", gin.Recovery())
```
例如这里g2的path就是/user，handlers就有Logger以及Recovery了。

RouterGroup还有其他的很多用法，例如Use函数可以为当前路由组添加中间件，Get、Post等函数可以添加路由。

## 中间件

## 路由树
