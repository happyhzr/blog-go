package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/insisthzr/blog-back/models"
)

func Signup(c *gin.Context) {
	user := &models.User{}
	err := c.Bind(user)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	err = user.Insert()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	user.Password = ""
	c.JSON(200, gin.H{"user": user})
}

func Login(c *gin.Context) {
	in := &models.User{}
	err := c.Bind(in)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	user, err := models.GetUserByName(in.Name)
	if err != nil {
		c.JSON(200, gin.H{"message": "user not exist"})
		return
	}
	if in.Password != user.Password {
		c.JSON(200, gin.H{"message": "password error"})
		return
	}
	user.Password = ""
	token := newJwtToken(jwt.MapClaims{"id": user.ID})
	c.JSON(200, gin.H{"user": user, "token": token})
}
