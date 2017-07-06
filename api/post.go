package api

import (
	"strconv"

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
	posts := busniess.ListPosts(in)
	return c.JSON(200, JSON{
		"posts": posts,
	})
}

func CountPosts(c echo.Context) error {
	count := busniess.CountPosts()
	return c.JSON(200, JSON{
		"count": count,
	})
}
