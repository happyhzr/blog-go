package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/insisthzr/blog-go/model"
	"github.com/spf13/viper"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Config  *config     `json:"config"`
}

type config struct {
	Page      int `json:"page" form:"page"`
	PageSize  int `json:"pageSize" form:"pageSize"`
	TotalPage int `json:"totalPage"`
}

func (c *config) setTotalPage(total int) {
	totalPage := total / c.PageSize
	if total%c.PageSize != 0 {
		totalPage++
	}
	c.TotalPage = totalPage
}

func newJwtToken(u *model.User) string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &myClaims{User: u})
	tokenStr, err := token.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		panic(err)
	}
	return tokenStr
}
