package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"

	"github.com/insisthzr/blog-back/models"
)

type myClaims struct {
	User *models.User `json:"user"`
	jwt.StandardClaims
}

func JwtAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequestWithClaims(
			c.Request,
			request.OAuth2Extractor,
			&myClaims{},
			func(token *jwt.Token) (interface{}, error) {
				b := ([]byte(secret))
				return b, nil
			})
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}
		claims, ok := token.Claims.(*myClaims)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}
		if claims.Valid() != nil {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}
		c.Set("user", claims.User)
		c.Next()
	}
}
