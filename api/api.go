package api

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/insisthzr/blog-go/api/controller"
)

func Start() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		user := v1.Group("/users")
		{
			user.POST("/signup", controller.Signup)
			user.POST("/login", controller.Login)
		}
		post := v1.Group("/posts")
		{
			post.GET("", controller.GetPosts)
			post.GET("/:id", controller.GetPost)
			post.POST("" /*, authMW*/, controller.CreatePost)
			//post.PATCH("/:id" /*, authMW*/, controller.UpdatePost)
			//post.DELETE("/:id" /*, authMW*/, controller.DeletePost)
		}
		//tag := v1.Group("/tags")
		{
			//tag.POST("" /*, authMW*/, controller.CreateTag)
			//tag.GET("", controller.ListTags)
			//tag.PATCH("", nil)
			//tag.DELETE("", nil)
		}
		category := v1.Group("/categorys")
		{
			category.GET("/:id", controller.GetCategory)
			category.POST("", controller.CreateCategory)
			//category.GET("", controller.ListCategorys)
			//category.PATCH("", nil)
			//category.DELETE("", nil)
		}
	}

	r.Run(":" + os.Getenv("port"))
}
