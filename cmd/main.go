package main

import (
	"github.com/Colocust/tinygo"
)

func main() {
	e := tinygo.Default()
	e.Run("127.0.0.1:6677")
}
