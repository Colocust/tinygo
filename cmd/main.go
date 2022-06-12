package main

import (
	"github.com/Colocust/tinygo"
	"github.com/Colocust/tinygo/controllers"
	"time"
)

func main() {
	e := tinygo.Default()
	userGroup := e.Group("user", tinygo.HandlerFuncWithTimeout(time.Millisecond*500), controllers.B, controllers.C)
	userGroup.Get("/info")
	e.Run("127.0.0.1:6677")
}
