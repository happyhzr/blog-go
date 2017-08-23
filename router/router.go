package router

import (
	"github.com/gin-gonic/gin"

	"github.com/insisthzr/blog-back/config"
	"github.com/insisthzr/blog-back/controller"
	"github.com/insisthzr/blog-back/util"
)

func Run() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	v1 := r.Group("/v1")

	user := v1.Group("/users")
	user.POST("/signup", controller.SignUp)
	user.POST("/signin", controller.SignIn)

	post := v1.Group("/posts")
	post.POST("", controller.CreatePost)
	post.GET("", controller.ListPosts)

	err := r.Run(config.HttpAddr)
	util.CheckError(err)
}
