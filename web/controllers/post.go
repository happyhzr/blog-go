package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/insisthzr/blog-back/models"
)

func ListPost(c *gin.Context) {
	o := c.Query("offset")
	if o == "" {
		o = "0"
	}
	l := c.Query("limit")
	if l == "" {
		l = "0"
	}
	offset, err := strconv.Atoi(o)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	limit, err := strconv.Atoi(l)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	if limit == 0 {
		limit = -1
	}

	posts, err := models.ListPostWithRange(offset, limit)
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
	post, err := models.GetPostByID(id)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, post)
}

type postIn struct {
	Title      string  `json:"title"`
	Content    string  `json:"content"`
	CategoryID int64   `json:"categoryID"`
	TagIDs     []int64 `json:"tagIDs"`
}

func CreatePost(c *gin.Context) {
	obj, _ := c.Get("user")
	user := obj.(*models.User)
	in := &postIn{}
	err := c.Bind(in)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	post := &models.Post{
		Title:   in.Title,
		Content: in.Content,
		Author:  user,
		Category: &models.Category{
			ID: in.CategoryID,
		},
		Tags: make([]*models.Tag, 0, len(in.TagIDs)),
	}
	for _, t := range in.TagIDs {
		post.Tags = append(post.Tags, &models.Tag{ID: t})
	}
	err = post.Create()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, post)
}

func UpdatePost(c *gin.Context) {
	obj, _ := c.Get("user")
	user := obj.(*models.User)
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	post := &models.Post{}
	err = c.Bind(post)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	post.ID = id
	post.AuthorID = user.ID
	err = post.Update()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "OK"})
}

func DeletePost(c *gin.Context) {
	obj, _ := c.Get("user")
	user := obj.(*models.User)
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "invalid id: " + idStr})
		return
	}
	post := &models.Post{ID: id, AuthorID: user.ID}
	err = post.Delete()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(400, gin.H{"message": "OK"})
}
