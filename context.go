package tinygo

import "net/http"

// 处理writer
type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
	Handlers HandlersChain
	index    int8
}

// index初始值为-1的原因是 如果为0后续++ 这样就不能认为控制ctx.Next
func (ctx *Context) Next() {
	ctx.index++
	for ctx.index < int8(len(ctx.Handlers)) {
		ctx.Handlers[ctx.index](ctx)
		ctx.index++
	}
}

func (ctx *Context) reset() {
	ctx.index = -1
}
