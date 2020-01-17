package controllers

import (
	"cocoyo/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type AuthController struct {}

func (auth *AuthController) Register(ctx *gin.Context) {
	var user models.User

	err := user.Create("cocoyo", "2430114823@qq.com", "123456")

	if err != nil {
		fmt.Println(err)
	}
}
