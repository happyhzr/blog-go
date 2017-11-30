package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/insisthzr/blog-back/models"
)

func CreateCategory(c *gin.Context) {
	category := &models.Category{}
	err := c.Bind(category)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	err = category.Create()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, category)
}
