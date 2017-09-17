package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/insisthzr/blog-back/models"
)

func CreatePost(c *gin.Context) {
	m, _ := c.Get("jwtMap")
	jwtMap := m.(jwt.MapClaims)
	post := &models.Post{}
	err := c.Bind(post)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	post.CreatedBy = &models.User{ID: int(jwtMap["id"].(float64))}
	err = post.Insert()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"post": post})
}

func ListPosts(c *gin.Context) {
	ol := offsetLimit{}
	err := c.Bind(&ol)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	if ol.Limit == 0 {
		ol.Limit = -1
	}
	count, err := models.CountPost()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	posts, err := models.ListPosts(ol.Offset, ol.Limit)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"posts": posts, "count": count})
}
