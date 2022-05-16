package routes

import (
	"github.com/gin-gonic/gin"
	"web_app/logger"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
