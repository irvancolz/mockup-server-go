package router

import (
	handler "mockup_server/handler"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GinRouter() *gin.Engine {
	gin.SetMode(os.Getenv("MODE"))
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":   "pong",
			"gin_mode":  os.Getenv("MODE"),
			"http_port": os.Getenv("ACTIVE_PORT"),
		})
	})

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", handler.Login)
		authRoutes.POST("/logout", handler.Logout)
	}

	return r
}
