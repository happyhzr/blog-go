package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func JwtAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(secret))
			return b, nil
		})
		if err != nil {
			c.AbortWithError(401, err)
		}
		m := token.Claims.(jwt.MapClaims)
		c.Set("jwtMap", m)
		c.Next()
	}
}
