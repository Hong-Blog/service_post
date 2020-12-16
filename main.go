package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	engine.GET("/post", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test sql",
		})
	})

	_ = engine.Run(":80")
}
