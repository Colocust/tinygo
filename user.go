package tinygo

import "log"

func GetUserInfo(c *Context) {
	log.Println(c.Request.RequestURI)
}