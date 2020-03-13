package main

import "github.com/gin-gonic/gin"

//go语言实现 websocket client

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("html/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "html/websocket.html", nil)
	})
	r.Run(":8080")
}
