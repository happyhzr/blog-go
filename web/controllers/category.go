package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/insisthzr/blog-back/services"
	"github.com/insisthzr/blog-back/utils"
)

func CreateCategory(c *gin.Context) {
	category := &services.Category{}
	err := c.Bind(category)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	err = category.Save()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"category": category})
}

func ListCategorys(c *gin.Context) {
	lo := utils.LimitOffset{}
	err := c.Bind(&lo)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	if lo.Limit == 0 {
		lo.Limit = -1
	}
	categorys, err := services.ListCategorys(lo.Limit, lo.Offset)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(400, gin.H{"categorys": categorys})
}
