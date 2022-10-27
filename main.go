package main

import (
	"github.com/PoteeDev/auth/middleware"
	"github.com/PoteeDev/scores/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", middleware.TokenAuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/scoreboard", middleware.TokenAuthMiddleware(), handlers.ShowScoreboard)
	r.GET("/user", middleware.TokenAuthMiddleware(), handlers.EntityScore)
	r.Run()
}
