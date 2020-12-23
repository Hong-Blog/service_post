package postHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service_post/models"
	"service_post/models/bizArticle"
	"strconv"
)

func GetArticleList(c *gin.Context) {
	var req bizArticle.GetArticleListRequest

	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	typeId, _ := strconv.Atoi(c.DefaultQuery("typeId", "0"))
	keyword := c.DefaultQuery("keyword", "")

	req.PageIndex = pageIndex
	req.PageSize = pageSize
	req.TypeId = typeId
	req.Keyword = keyword

	list, count := bizArticle.GetArticleList(req)

	var response models.PagedResponse
	response.Data = list
	response.Total = count
	c.JSON(http.StatusOK, response)
}