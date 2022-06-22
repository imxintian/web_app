package routes

import (
	"github.com/gin-gonic/gin"
	"time"
	"web_app/controller"
	"web_app/logger"
	"web_app/middlewares"
)

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1")

	// 注册业务路由
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware(), middlewares.RateLimitMiddleware(2*time.Second, 1)) //api 限流
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/post/", controller.GetPostListHandler)
		v1.GET("post2/", controller.GetPostList2Handler)
		v1.POST("/vote", controller.PostVoteHandler)

	}

	return r
}
