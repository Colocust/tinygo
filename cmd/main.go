package main

import (
	"github.com/Colocust/tinygo"
	"github.com/Colocust/tinygo/controllers"
)

func main() {
	e := tinygo.Default()
	userGroup := e.Group("user", controllers.GetUserInfo, controllers.SetUserInfo)
	userGroup.Get("/info")
	e.Run("127.0.0.1:6677")
}
