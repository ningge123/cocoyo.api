package routes

import (
	"cocoyo/http/controllers"
	"cocoyo/pkg/cross"
	"cocoyo/pkg/e"
	"cocoyo/pkg/jwt"
	"cocoyo/pkg/setting"
	"cocoyo/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(
		gin.Logger(),
		gin.Recovery(),
		cross.Handler(),
		e.Handler(),
	)

	r.StaticFS(setting.Filesystem.Key("root").String(), http.Dir(util.GetAppAbsolutePath() + setting.Filesystem.Key("root").String()))

	api := r.Group("/api")
	{
		// 认证
		var authController *controllers.Auth
		auth := api.Group("/auth")
		{
			auth.POST("/login", authController.Login)
			auth.POST("/register", authController.Register)
		}

		// 用户
		var userController *controllers.User
		var fileController *controllers.File
		api.POST("/user/exists", userController.Exists)
		authorized := api.Use(jwt.TokenAuthMiddleware)
		{
			authorized.GET("/me", userController.Me)
			authorized.POST("/user/update", userController.Update)
			authorized.POST("/files/upload", fileController.Upload)
		}

		// 网盘
		//var cloud *controllers.Cloud
		//api.GET("/files", cloud.Index)
		//api.POST("/files", cloud.Update)
		//api.DELETE("/files", cloud.Delete)
		//api.POST("/files/upload", cloud.Upload)
		//api.GET("/files/download", cloud.Download)
	}

	return r
}