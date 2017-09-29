package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/insisthzr/blog-back/services"
)

func Signup(c *gin.Context) {
	user := &services.User{}
	err := c.Bind(user)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	err = user.Signup()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"user": user})
}

func Login(c *gin.Context) {
	user := &services.User{}
	err := c.Bind(user)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	err = user.Login()
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	token := newJwtToken(jwt.MapClaims{"id": user.ID})
	c.JSON(200, gin.H{"user": user, "token": token})
}
