package tinygo

import (
	"github.com/Colocust/tinygo/render"
	"math"
	"net/http"
	"sync"
)

// 处理writer
type Context struct {
	Request *http.Request
	Writer  ResponseWriter

	engine *Engine

	Handlers HandlersChain
	index    int8

	hasTimeout bool
	mutex      sync.Mutex
}

const AbortIndex = math.MaxInt8 >> 1

// index初始值为-1的原因是 如果为0后续++ 这样就不能人为控制ctx.Next
func (ctx *Context) Next() {
	ctx.index++
	for ctx.index < int8(len(ctx.Handlers)) {
		ctx.Handlers[ctx.index](ctx)
		ctx.next()
	}
}

func (ctx *Context) next() {
	ctx.index++
}

func (ctx *Context) Abort() {
	ctx.index = AbortIndex
}

func (ctx *Context) reset() {
	ctx.index = -1
}

func (ctx *Context) Json(status int, data interface{}) {
	ctx.render(status, render.Json{
		Data: data,
	})
}

func (ctx *Context) render(status int, r render.Render) {
	if !ctx.HasTimeout() {
		ctx.status(status)
		r.Render(ctx.Writer)
	}
}

func (ctx *Context) status(status int) {
	ctx.Writer.WriteHeader(status)
}

func (ctx *Context) SetTimeout() {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()

	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()

	return ctx.hasTimeout
}
