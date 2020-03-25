package controllers

import (
	"cocoyo/http/requests"
	"cocoyo/models"
	"cocoyo/pkg/jwt"
	"cocoyo/pkg/response"
	"cocoyo/pkg/setting"
	"cocoyo/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/o1egl/govatar"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Auth struct {}

func (auth *Auth) Login(ctx *gin.Context)  {
	valid := validation.Validation{}
	register := requests.Login{Email: ctx.PostForm("email"), Password: ctx.PostForm("password")}

	b, err := valid.Valid(&register)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ParameterError(err.Error()))

		return
	}

	if !b {
		for _, err := range valid.Errors {
			ctx.JSON(http.StatusUnprocessableEntity, response.ParameterError(err.Error()))

			return
		}
	}

	var user models.User
	// 查询数据库 验证是否有这个用户
	if models.ReturnDB().Where("email = ?", register.Email).First(&user).RecordNotFound() {
		ctx.JSON(http.StatusUnprocessableEntity, response.ParameterError("用户名或密码错误"))

		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(register.Password))

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ParameterError("用户名或密码错误"))

		return
	}

	// 生成jwt
	token, err := jwt.GenerateToken(&user)

	if err != nil {
		panic(err)
	}

	data := make(map[string]string)
	data["access_token"] = token

	ctx.JSON(http.StatusOK, response.Response(data))
}

func (auth *Auth) Register(ctx *gin.Context)  {
	valid := validation.Validation{}
	register := requests.Register{Username:ctx.PostForm("username"), Email: ctx.PostForm("email"), Password: ctx.PostForm("password")}

	b, err := valid.Valid(&register)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ParameterError(err.Error()))

		return
	}

	if !b {
		for _, err := range valid.Errors {
			ctx.JSON(http.StatusUnprocessableEntity, response.ParameterError(err.Error()))

			return
		}
	}

	avatarPath := setting.Filesystem.Key("root").String()
	// 创建文件夹
	util.MakeDir(avatarPath)
	password, _ := bcrypt.GenerateFromPassword([]byte(ctx.PostForm("password")), bcrypt.DefaultCost)
	// 处理
	user := models.User{
		Email: ctx.PostForm("email"),
		Password: string(password),
		Username: ctx.PostForm("username"),
		Avatar:  "/avatar" + util.Md5(ctx.PostForm("email")) + ".jpg",
	}

	// 创建用户
	models.ReturnDB().Select("email", "password", "username", "avatar").Create(&user)

	// 生成头像
	govatar.GenerateFileForUsername(govatar.MALE, user.Username, util.GetAppAbsolutePath() + user.Avatar)
	models.ReturnDB().Where("email = ?", user.Email).First(&user)

	token, _ := jwt.GenerateToken(&user)
	data := make(map[string]string)
	data["access_token"] = token

	ctx.JSON(http.StatusOK, response.Response(data))
}