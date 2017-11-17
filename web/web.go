package web

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/insisthzr/blog-back/config"
	"github.com/insisthzr/blog-back/utils"
	"github.com/insisthzr/blog-back/web/controllers"
)

func Start() {
	//authMW := middlewares.JwtAuth(config.DefaultConfig.Jwt.Secret)

	r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
			return
		})
		user := v1.Group("/users")
		{
			user.POST("/signup", controllers.Signup)
			user.POST("/login", controllers.Login)
		}
		post := v1.Group("/posts")
		{
			post.GET("", controllers.ListPosts)
			post.GET("/:id", controllers.GetPost)
			post.POST("", controllers.CreatePost)
			post.PATCH("/:id", controllers.UpdatePost)
			post.DELETE("/:id", controllers.DeletePost)
		}
		tag := v1.Group("/tags")
		{
			tag.POST("", controllers.CreateTag)
			tag.GET("", controllers.ListTags)
			tag.PATCH("", nil)
			tag.DELETE("", nil)
		}
		category := v1.Group("/categorys")
		{
			category.POST("", controllers.CreateCategory)
			category.GET("", controllers.ListCategorys)
			category.PATCH("", nil)
			category.DELETE("", nil)
		}
	}

	srv := &http.Server{
		Addr:    config.DefaultConfig.Http.Addr,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			utils.Sugar.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	utils.Sugar.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		utils.Sugar.Fatal("Server Shutdown:", err)
	}
	utils.Sugar.Info("Server exiting")
}
