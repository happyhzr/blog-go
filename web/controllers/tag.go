package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/insisthzr/blog-back/services"
	"github.com/insisthzr/blog-back/utils"
)

func CreateTag(c *gin.Context) {
	tag := &services.Tag{}
	err := c.Bind(tag)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	err = tag.Save()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"tag": tag})
}

func ListTags(c *gin.Context) {
	lo := utils.LimitOffset{}
	err := c.Bind(&lo)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	if lo.Limit == 0 {
		lo.Limit = -1
	}
	tags, err := services.ListTags(lo.Limit, lo.Offset)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(400, gin.H{"tags": tags})
}
