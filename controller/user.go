package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/insisthzr/blog-back/model"
)

type UserIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (i *UserIn) NewModel() *model.User {
	return &model.User{
		Username: i.Username,
		Password: i.Password,
	}
}

type UserOut struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

func NewUserOut(m *model.User) *UserOut {
	return &UserOut{
		Id:       m.Id,
		Username: m.Username,
	}
}

func SignUp(c *gin.Context) {
	in := &UserIn{}
	err := c.BindJSON(in)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	user := &model.User{
		Username: in.Username,
		Password: in.Password,
	}
	user.Insert()
	out := NewUserOut(user)
	c.JSON(200, out)
}

func SignIn(c *gin.Context) {
	in := &UserIn{}
	err := c.BindJSON(in)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	user := model.GetUserByUsername(in.Username)
	if user == nil {
		c.JSON(400, gin.H{"message": "username or password error"})
		return
	}
	if in.Password != user.Password {
		c.JSON(400, gin.H{"message": "username or password error"})
		return
	}
	out := NewUserOut(user)
	c.JSON(200, out)
}
