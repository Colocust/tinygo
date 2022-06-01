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

## 中间件

## 路由树
