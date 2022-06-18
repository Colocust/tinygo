package controllers

import (
	"github.com/Colocust/tinygo"
	"log"
	"time"
)

func A(ctx *tinygo.Context) {
	log.Println("A Start")

	log.Println("A End")
}

func B(ctx *tinygo.Context) {
	log.Println("B Start")

	time.Sleep(time.Second * 1)

	ctx.Json(200, "504 Gateway Time-out")

	log.Println("B End")
}

func C(ctx *tinygo.Context) {
	log.Println("C Start")
	//time.Sleep(time.Second * 1)
	ctx.Json(200, map[string]interface{}{
		"data": 200,
	})
	log.Println("C End")
}
