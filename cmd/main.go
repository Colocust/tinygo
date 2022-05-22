package main

import (
	"fmt"
	"github.com/Colocust/tinygo"
)

func main() {
	var e tinygo.IRoutes
	e = &tinygo.Engine{

	}
	fmt.Println(e)
	e.Use()
}
