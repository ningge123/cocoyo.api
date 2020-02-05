package routers

import (
	"cocoyo/controllers"
	"cocoyo/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cross())

	api := r.Group("/api")
	{
		// 节点
		var node *controllers.NodeController
		api.GET("/nodes", node.Index)
		// 用户
		var user *controllers.UserController
		api.GET("/users", user.Index)
		api.POST("/user/exists", user.Exists)
		// 注册
		var auth *controllers.AuthController
		api.POST("/register", auth.Register)
	}

	return r
}