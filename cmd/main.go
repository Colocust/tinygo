package main

import "github.com/Colocust/tinygo"

func main() {
	engine := tinygo.Default()
	engine.NoRoute(func(ctx *tinygo.Context) {
		ctx.JSON(405, tinygo.H{
			"data": "404 not found",
		})
	})
	engine.Run("127.0.0.1:7878")
}
