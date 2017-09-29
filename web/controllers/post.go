package controllers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/insisthzr/blog-back/models"
	"github.com/insisthzr/blog-back/services"
)

type postIn struct {
	models.PostSelector
}

func ListPosts(c *gin.Context) {
	in := postIn{}
	err := c.Bind(&in)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	if in.Limit == 0 {
		in.Limit = -1
	}
	posts, err := services.ListPostsS(in.PostSelector)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"posts": posts})
}

func GetPost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "invalid id: " + idStr})
		return
	}
	p, err := services.GetPost(id)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"post": p})
}

func CreatePost(c *gin.Context) {
	//m, _ := c.Get("jwtMap")
	//jwtMap := m.(jwt.MapClaims)
	post := &services.Post{}
	err := c.Bind(post)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	//post.AuthorID = int64(jwtMap["id"].(float64))
	post.AuthorID = 1
	post.CreatedAt = time.Now().Unix()
	err = post.Save()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"post": post})
}

func UpdatePost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "invalid id: " + idStr})
		return
	}
	post := &services.Post{}
	err = c.Bind(post)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	err = post.Update(id)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(400, gin.H{"post": post})
}

func DeletePost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "invalid id: " + idStr})
		return
	}
	err = services.DeletePost(id)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(400, gin.H{"message": "OK"})
}
