package controllers

import (
	"cocoyo/models"
	"cocoyo/pkg/util/response"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {}

func (c *UserController) Index(ctx *gin.Context)  {
	test := models.ExtendsFields{Company:"asdd", Location:"asd", HomeUrl:"123", Github:"asd", Twitter:"qwe", Facebook:"sad", Instagram:"aswqe", Telegram:"qwe", Coding:"qwe", Steam:"", Weibo:""}
	testJson, _ := json.Marshal(test)
	var a models.ExtendsFields
	json.Unmarshal(testJson, &a)
	fmt.Println(a.Company)
	fmt.Println(string(testJson))
}

func (c *UserController) Exists(ctx *gin.Context) {
	email    := ctx.PostForm("email")

	fmt.Println(email)
	// 验证器
	validator := struct {
		Email    string `valid:"Required; MaxSize(50); Email"`
	}{Email:email}

	valid := validation.Validation{}

	ok, _ := valid.Valid(&validator)

	if ok {
		var user models.User
		user.Email = email

		err := user.BeforeCreate()

		if err != nil {
			ctx.JSON(http.StatusOK, response.ReturnMsgFunc(response.Code{Code: response.ERROR, Message: err.Error()}, nil))
		} else {
			ctx.JSON(http.StatusOK, response.ReturnMsgFunc(response.Code200, map[string]bool{"success": true}))
		}
	} else {

		firstErrorMsg := valid.Errors[0]

		ctx.JSON(http.StatusOK, response.ReturnMsgFunc(response.Code{Code: response.ERROR, Message: firstErrorMsg.Message}, nil))
	}
}