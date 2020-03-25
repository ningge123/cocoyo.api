package controllers

import (
	"cocoyo/models"
	"cocoyo/pkg/jwt"
	"cocoyo/pkg/response"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type User struct {}

// 注册 判断用户是否已存在
func (u *User) Exists(ctx *gin.Context) {
	username := ctx.PostForm("username")
	email    := ctx.PostForm("email")

	fmt.Println(username, email)

	if username != "" || email != "" {
		var user models.User
		data := make(map[string]int)
		data["success"] = 0

		if models.ReturnDB().Where("email = ?", email).Or("username = ?", username).First(&user).RecordNotFound() {
			data["success"] = 1
		}

		ctx.JSON(http.StatusOK, response.Response(data))

		return
	}

	ctx.JSON(http.StatusUnprocessableEntity, response.ParameterError("参数不能为空"))
}

// 获取获取信息
func (u *User) Me(ctx *gin.Context) {
	token := ctx.GetString(jwt.ContextKeyUserObj)

	tokenUser, _ := jwt.ParseToken(token)

	var user models.User

	models.ReturnDB().Where("id = ?", tokenUser.Id).First(&user)

	ctx.JSON(http.StatusOK, response.Response(user))
}

func (u *User) Update(ctx *gin.Context)  {
	token := ctx.GetString(jwt.ContextKeyUserObj)

	currentUser, err := jwt.ParseToken(token)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.Authorization())

		return
	}

	userId := ctx.PostForm("user_id")
	intUserId, _ := strconv.Atoi(userId)

	if currentUser.Id != intUserId {
		ctx.JSON(http.StatusUnprocessableEntity, response.ParameterError("你没有权限修改"))

		return
	}

	var user models.User
	var extend models.ExtendsFields

	extend.Github = ctx.PostForm("github")
	extend.HomeUrl = ctx.PostForm("home_url")
	extend.Weibo = ctx.PostForm("wei_bo")

	extendByte, _ := json.Marshal(extend)

	// 更新用户信息
	err = models.ReturnDB().Model(&user).Where("id = ?", userId).Updates(models.User{
		Avatar: ctx.PostForm("avatar"),
		Bio: ctx.PostForm("bio"),
		Gender: ctx.PostForm("gender"),
		Extends: string(extendByte),
	}).Error

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ParameterError(err.Error()))

		return
	}

	ctx.JSON(http.StatusOK, response.SuccessNotContent())
}