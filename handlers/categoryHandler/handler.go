package categoryHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service_post/models"
	"service_post/models/bizType"
	"service_post/validator"
)

func CategoryList(c *gin.Context) {
	typeList := bizType.GetTypeList()
	c.JSON(http.StatusOK, typeList)
}

func AddBizType(c *gin.Context) {
	var req bizType.AddBizTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: validator.Translate(err)})
		return
	}

	if err := bizType.AddBizType(req); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.String(http.StatusOK, "")
}
