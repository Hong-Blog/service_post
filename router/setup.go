package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service_post/handlers/categoryHandler"
	"service_post/handlers/postHandler"
	"service_post/handlers/tagHandler"
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

	tags := engine.Group("/tags")
	{
		tags.GET("", tagHandler.TagList)
		tags.GET("/", tagHandler.TagList)
		tags.POST("", tagHandler.AddBizTag)
		tags.GET("/:id", tagHandler.GetById)
		tags.PUT("/:id", tagHandler.UpdateById)
		tags.DELETE("/:id", tagHandler.DeleteById)
	}

	posts := engine.Group("/posts")
	{
		posts.GET("", postHandler.GetArticleList)
		posts.GET("/", postHandler.GetArticleList)
		posts.POST("", postHandler.AddArticle)
		posts.GET("/:id", postHandler.GetDetailById)
		posts.PUT("/:id", postHandler.UpdateById)
		posts.DELETE("/:id", postHandler.DeleteById)
	}
}
