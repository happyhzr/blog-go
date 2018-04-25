package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/insisthzr/blog-back/model"
)

func CreateCategory(c *gin.Context) {
	category := &model.Category{}
	err := c.Bind(category)
	if err != nil {
		c.JSON(200, response{Code: 1, Message: err.Error()})
		return
	}
	err = category.Create()
	if err != nil {
		c.JSON(200, response{Code: 1, Message: err.Error()})
		return
	}
	c.JSON(200, response{Code: 0, Data: category})
}

func GetCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(200, response{Code: 1, Message: err.Error()})
		return
	}
	category, err := model.GetCategoryByID(id)
	if err != nil {
		c.JSON(200, response{Code: 1, Message: err.Error()})
		return
	}
	c.JSON(200, response{Code: 0, Data: category})
}
