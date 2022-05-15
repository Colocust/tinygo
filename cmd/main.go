package main

import "github.com/Colocust/tinygo"

func main() {
	engine := tinygo.Default()
	engine.NoRoute(func(ctx *tinygo.Context) {
		ctx.JSON(405,"405 no found")
	})
	engine.NoRoute(func(ctx *tinygo.Context) {
		ctx.JSON(406,"405 no found")
	})
	engine.Run("127.0.0.1:7878")
}
