package app

type (
	HandlerFunc   func(ctx *Context)
	HandlersChain []HandlerFunc
)
