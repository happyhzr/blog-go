package web

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"

	"github.com/insisthzr/blog-back/config"
	"github.com/insisthzr/blog-back/web/controllers"
	"github.com/insisthzr/blog-back/web/middlewares"
)

func Start() {
	router := gin.Default()

	api := router.Group("/api")

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
			post.Use(middlewares.JwtAuth(config.DefaultConfig.Jwt.Secret)).POST("", controllers.CreatePost)
		}
	}

	srv := &http.Server{
		Addr:    config.DefaultConfig.Http.Addr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logrus.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("Server Shutdown:", err)
	}
	logrus.Info("Server exiting")
}
