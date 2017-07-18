package api

import (
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	"github.com/insisthzr/blog-back/busniess"
	"github.com/insisthzr/blog-back/util"
)

func CreatePost(c echo.Context) error {
	userID := util.GetIDFromJWT(c.Get("user").(*jwt.Token))
	postIn := &busniess.PostIn{}
	err := c.Bind(postIn)
	if err != nil {
		return err
	}
	postIn.CreatedBy = userID
	in := &busniess.CreatePostIn{PostIn: *postIn}
	out := busniess.CreatePost(in)
	logrus.WithFields(logrus.Fields{"out": out}).Infoln("CreatePost")
	return c.JSON(200, JSON{"post": out})
}

func ListPosts(c echo.Context) error {
	skipStr := c.QueryParam("skip")
	limitStr := c.QueryParam("limit")
	skip, _ := strconv.Atoi(skipStr)
	limit, _ := strconv.Atoi(limitStr)
	in := &busniess.ListPostsIn{
		Range: busniess.SkipLimit{
			Skip:  skip,
			Limit: limit,
		},
	}
	out := busniess.ListPosts(in)
	logrus.WithFields(logrus.Fields{"out": out}).Infoln("ListPosts")
	return c.JSON(200, JSON{"posts": out})
}

func CountPosts(c echo.Context) error {
	out := busniess.CountPosts()
	logrus.WithFields(logrus.Fields{"out": out}).Infoln("CountPosts")
	return c.JSON(200, JSON{"count": out})
}
