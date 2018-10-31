package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/insisthzr/blog-go/model"
)

func GetPosts(c *gin.Context) {
	cfg := &config{}
	err := c.Bind(cfg)
	if err != nil {
		c.JSON(200, response{Code: 1, Message: err.Error()})
		return
	}
	posts, total, err := model.GetPosts(cfg.Page, cfg.PageSize)
	if err != nil {
		c.JSON(200, response{Code: 1, Message: err.Error()})
		return
	}
	cfg.setTotalPage(total)
	c.JSON(200, response{Code: 0, Data: posts, Config: cfg})
}

type getPostIn struct {
	ID int64 `form:"id"`
}

func GetPost(c *gin.Context) {
	in := &getPostIn{}
	err := c.Bind(in)
	if err != nil {
		c.JSON(200, response{Code: 1, Message: err.Error()})
		return
	}
	post, err := model.GetPostByID(in.ID)
	if err != nil {
		c.JSON(200, response{Code: 1, Message: err.Error()})
		return
	}
	c.JSON(200, response{Code: 0, Data: post})
}

func CreatePost(c *gin.Context) {
	post := &model.Post{}
	err := c.Bind(post)
	if err != nil {
		c.JSON(200, response{Code: 1, Message: err.Error()})
		return
	}
	err = post.Create()
	if err != nil {
		c.JSON(200, response{Code: 1, Message: err.Error()})
		return
	}
	c.JSON(200, response{Code: 0, Data: post})
}
