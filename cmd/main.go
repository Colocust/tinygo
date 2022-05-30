package main

import (
	"github.com/Colocust/tinygo"
)

func main() {
	e := tinygo.Default()
	userGroup := e.Group("user", tinygo.GetUserInfo, tinygo.SetUserInfo)
	userGroup.Get("/info")
	e.Run("127.0.0.1:6677")
}
