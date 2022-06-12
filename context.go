package tinygo

import (
	"github.com/Colocust/tinygo/render"
	"math"
	"net/http"
	"sync"
)

// 处理writer
type Context struct {
	Request    *http.Request
	Response   http.ResponseWriter
	Handlers   HandlersChain
	index      int8
	mutex      sync.Mutex
	hasTimeout bool
}

const AbortIndex = math.MaxInt8 >> 1

// index初始值为-1的原因是 如果为0后续++ 这样就不能人为控制ctx.Next
func (ctx *Context) Next() {
	ctx.index++
	for ctx.index < int8(len(ctx.Handlers)) {
		ctx.Handlers[ctx.index](ctx)
		ctx.index++
	}
}

func (ctx *Context) Abort() {
	ctx.index = AbortIndex
}

func (ctx *Context) reset() {
	ctx.index = -1
}

func (ctx *Context) Json(status int, data interface{}) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	if ctx.HasTimeout() {
		return
	}
	ctx.render(status, render.Json{
		Data: data,
	})
}

func (ctx *Context) render(status int, r render.Render) {
	ctx.status(status)
	r.Render(ctx.Response)
}

func (ctx *Context) status(status int) {
	ctx.Response.WriteHeader(status)
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

func (ctx *Context) SetTimeout() {
	ctx.hasTimeout = true
}
