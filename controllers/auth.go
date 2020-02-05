package controllers

import (
	"cocoyo/models"
	"cocoyo/pkg/e"
	"cocoyo/pkg/util/response"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {}

func (auth *AuthController) Register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	email    := ctx.PostForm("email")

	// 验证器
	validator := struct {
		Username string `valid:"Required; MaxSize(50)"`
		Password string `valid:"Required; MaxSize(12); MinSize(6)"`
		Email    string `valid:"Required; MaxSize(50); Email"`
	}{Username:username, Password:password, Email:email}

	valid := validation.Validation{}
	ok, _ := valid.Valid(&validator)

	if ok {
		var user models.User

		token, err := user.Create(username, email, password)

		if err != nil {
			ctx.JSON(http.StatusOK, response.ReturnMsgFunc(response.Code{Code: response.ERROR, Message: err.Error()}, nil))
		} else {
			ctx.JSON(http.StatusOK, response.ReturnMsgFunc(response.Code200, map[string]string{"token" : token}))
		}
	} else {
		for _, v := range valid.Errors {
			e.New(e.WARNING, v.Key + v.Message)
		}
		ctx.Status(http.StatusUnauthorized)
	}
}
