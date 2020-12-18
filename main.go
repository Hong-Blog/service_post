package main

import (
	"github.com/gin-gonic/gin"
	"service_post/router"
)

func main() {
	engine := gin.Default()

	router.SetupRouter(engine)

	addr := ":80"
	if gin.Mode() == gin.DebugMode {
		addr = ":8083"
	}
	_ = engine.Run(addr)
}
