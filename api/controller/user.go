package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/insisthzr/blog-back/model"
)

func Signup(c *gin.Context) {
	user := &model.User{}
	err := c.Bind(user)
	if err != nil {
		c.JSON(200, response{Code: 1, Message: err.Error()})
		return
	}
	err = user.Signup()
	if err != nil {
		c.JSON(200, response{Code: 1, Message: err.Error()})
		return
	}
	c.JSON(200, response{Code: 0, Data: user})
}

func Login(c *gin.Context) {
	user := &model.User{}
	err := c.Bind(&user)
	if err != nil {
		c.JSON(200, response{Code: 1, Message: err.Error()})
		return
	}
	err = user.Login()
	if err != nil {
		c.JSON(200, response{Code: 1, Message: err.Error()})
		return
	}
	user.Password = ""
	c.JSON(200, response{Code: 0, Data: user})
}
