package controller

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/insisthzr/blog-back/model"
)

type PostIn struct {
	Title   string
	Content string
}

func (i *PostIn) NewModel() *model.Post {
	now := time.Now().Unix()
	return &model.Post{
		Title:     i.Title,
		Content:   i.Content,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

type PostOut struct {
	Id        int64
	Title     string
	Content   string
	CreatedAt int64
	UpdatedAt int64
}

func NewPostOut(m *model.Post) *PostOut {
	return &PostOut{
		Id:        m.Id,
		Title:     m.Title,
		Content:   m.Content,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func CreatePost(c *gin.Context) {
	in := &PostIn{}
	err := c.BindJSON(in)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	post := in.NewModel()
	post.Insert()
	out := NewPostOut(post)
	c.JSON(200, out)
}

func ListPosts(c *gin.Context) {
	posts := model.ListPosts()
	outs := make([]*PostOut, 0, len(posts))
	for _, post := range posts {
		outs = append(outs, NewPostOut(post))
	}
	c.JSON(200, outs)
}
