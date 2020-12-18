package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"service_post/db"
	"service_post/models/bizArticle"
)

func main() {
	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	engine.GET("/post", func(c *gin.Context) {

		sql := `
select *
from biz_article
`
		var list []bizArticle.Article
		if err := db.Db.Select(&list, sql); err != nil {
			log.Panicln(err.Error())
		}

		c.JSON(http.StatusOK, list)
	})

	addr := ":80"
	if gin.Mode() == gin.DebugMode {
		addr = ":8083"
	}
	_ = engine.Run(addr)
}
