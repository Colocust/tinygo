package main

import (
	"github.com/Colocust/tinygo"
	"github.com/Colocust/tinygo/controllers"
)

func main() {
	tg := tinygo.Default()

	userGroup := tg.Group("user", controllers.B, controllers.C)
	userGroup.Get("/info")

	if err := tg.Run("127.0.0.1:6677"); err != nil {
		panic(err)
	}
}
