package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service_post/handlers/categoryHandler"
)

func SetupRouter(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	categories := engine.Group("/categories")
	{
		categories.GET("", categoryHandler.CategoryList)
		categories.GET("/", categoryHandler.CategoryList)
		categories.POST("", categoryHandler.AddBizType)
		categories.GET("/:id", categoryHandler.GetById)
		categories.PUT("/:id", categoryHandler.UpdateById)
		categories.DELETE("/:id", categoryHandler.DeleteById)
	}
}
