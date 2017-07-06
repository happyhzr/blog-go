package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/insisthzr/blog-back/conf"
)

func NewJWTToken(id string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Duration(conf.JWTExp) * time.Second).Unix()
	t, err := token.SignedString(conf.JWTKey)
	if err != nil {
		panic(err)
	}
	return t
}

func GetIDFromJWT(user *jwt.Token) string {
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}
