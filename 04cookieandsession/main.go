package main

import (
	"code.oldbody.com/studygolang/mylearn/04cookieandsession/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Any("/login", handlers.LoginHandler)
	r.GET("/index", handlers.IndexHandler)
	r.GET("/vip", handlers.CookieMiddleware, handlers.VipHandler)
	r.Run(":8081")
}
