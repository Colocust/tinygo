package tinygo

import "log"

func GetUserInfo(c *Context) {
	log.Println("GetUserInfoStart")
	c.Next()
	log.Println("GetUserInfoEnd")
}

func SetUserInfo(c *Context) {
	log.Println("SetUserInfoStart")
	c.Json(200, map[string]interface{}{
		"code": 200,
		"data": map[string]interface{}{
			"info": "ss",
		},
	})
	log.Println("SetUserInfoEnd")
}
