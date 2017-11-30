package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/insisthzr/blog-back/models"
)

func CreateTag(c *gin.Context) {
	tag := &models.Tag{}
	err := c.Bind(tag)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	err = tag.Create()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, tag)
}
