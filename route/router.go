package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/insisthzr/blog-back/api"
	"github.com/insisthzr/blog-back/conf"
)

func Init() *echo.Echo {

	e := echo.New()
	e.Debug = true

	// Set Bundle MiddleWare
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())

	// Set Custom MiddleWare

	// Routes
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	v1 := e.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			user.POST("/login", api.Login)
			user.POST("/signup", api.Signup)
		}
		post := v1.Group("/posts")
		{
			post.POST("", api.CreatePost, middleware.JWT(conf.JWTKey))
			post.GET("", api.ListPosts)
			post.GET("/count", api.CountPosts)
		}
		archive := v1.Group("/archives")
		{
			archive.GET("", api.ListArchives)
		}
		file := v1.Group("/files")
		{
			file.POST("", api.UploadFile, middleware.JWT(conf.JWTKey))
		}
	}
	return e
}
