package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/insisthzr/blog-back/config"
)

func getIDFromJWT(token *jwt.Token) int {
	id := int(token.Claims.(jwt.MapClaims)["id"].(float64))
	return id
}

func newJwtToken(m jwt.MapClaims) string {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = m
	tokenStr, err := token.SignedString([]byte(config.DefaultConfig.Jwt.Secret))
	if err != nil {
		panic(err)
	}
	return tokenStr
}
