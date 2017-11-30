package controllers

import (
	"github.com/dgrijalva/jwt-go"

	"github.com/insisthzr/blog-back/config"
	"github.com/insisthzr/blog-back/models"
)

type offsetLimit struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func newJwtToken(u *models.User) string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &myClaims{User: u})
	tokenStr, err := token.SignedString([]byte(config.C.Jwt.Secret))
	if err != nil {
		panic(err)
	}
	return tokenStr
}
