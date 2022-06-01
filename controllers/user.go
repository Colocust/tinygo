package controllers

import (
	"github.com/Colocust/tinygo"
	"log"
)

func GetUserInfo(c *tinygo.Context) {
	log.Println("GetUserInfoStart")
	c.Next()
	log.Println("GetUserInfoEnd")
}

func SetUserInfo(c *tinygo.Context) {
	log.Println("SetUserInfoStart")
	c.Json(200, map[string]interface{}{
		"code": 200,
		"data": map[string]interface{}{
			"info": "ss",
		},
	})
	log.Println("SetUserInfoEnd")
}
