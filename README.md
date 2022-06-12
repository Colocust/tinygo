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

## Context


## 中间件

### 结构
```Golang
type HandlerFunc func(*Context)
```
中间件，其原理就是对一个方法进行包裹装饰，然后返回同类型的方法。

应用场景大多是需要对某一类函数进行通用的前置或者后置处理

### 具体使用
当框架匹配到具体路由并获取到需要执行的中间件后，会调用Next函数按照中间件添加的顺序依次执行。

```go
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}
```
例如此时我有以下三个中间件，并按顺序依次绑定在指定路由中：
```go
func A(ctx *gin.Context) {
	log.Println("A Start")
	log.Println("A End")
}

func B(ctx *gin.Context) {
	log.Println("B Start")
	log.Println("B End")
}

func C(ctx *gin.Context) {
	log.Println("C Start")
	log.Println("C End")
}
```
如上述逻辑所示，每个中间件都会完成自己所有的业务逻辑后才会开始执行下一个中间件。
```
A Start
A End
B Start
B End
C Start
C End
```
但有些业务场景需要我们在结束之前就开始调用下一个中间件，例如超时控制中间件。

那么此时我们就需要在中间件里手动调用一下Next函数来改变其执行过程了。还是刚刚的例子：
```go
func A(ctx *gin.Context) {
	log.Println("A Start")
	ctx.Next()
	log.Println("A End")
}

func B(ctx *gin.Context) {
	log.Println("B Start")
	log.Println("B End")
}

func C(ctx *gin.Context) {
	log.Println("C Start")
	log.Println("C End")
}
```
此时日志打印结果便为：
```
A Start
B Start
B End
C Start
C End
A End
```
## 路由树
