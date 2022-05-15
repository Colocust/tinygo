package tinygo

import (
	"github.com/Colocust/tinygo/render"
	"net/http"
)

type Context struct {
	Request  *http.Request
	Writer   *ResponseWriter
	engine   *Engine
	index    int8
	handlers HandlersChain
}

func (ctx *Context) Reset() {
	ctx.handlers = nil
	ctx.index = -1
}

func (ctx *Context) Next() {
	ctx.index++
	for ctx.index < int8(len(ctx.handlers)) {
		ctx.handlers[ctx.index](ctx)
		ctx.index++
	}
}

func (ctx *Context) JSON(status int, data any) {
	ctx.Render(status, &render.Json{Data: data})
}

func (ctx *Context) Render(status int, render render.Render) {
	ctx.status(status)
	if err := render.Render(ctx.Writer); err != nil {
		panic(err)
	}
}

func (ctx *Context) status(status int) {
	ctx.Writer.status = status
}
