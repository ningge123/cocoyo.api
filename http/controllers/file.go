package controllers

import (
	"cocoyo/pkg/response"
	"cocoyo/pkg/setting"
	"cocoyo/pkg/util"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type File struct {}

func (f *File) Upload(ctx *gin.Context)  {
	responseMap := make(map[string]string)
	fileMeta, err := ctx.FormFile("file")

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ParameterError(err.Error()))

		return
	}

	file, err := fileMeta.Open()

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ParameterError(err.Error()))

		return
	}

	defer file.Close()

	responseMap["url"] = setting.LoadFilesystem().Key("root").String() + "/avatar/" + fileMeta.Filename

	newFile, err := os.Create(util.GetAppAbsolutePath() + responseMap["url"])

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ParameterError(err.Error()))

		return
	}

	_, err = io.Copy(newFile, file)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ParameterError(err.Error()))

		return
	}

	ctx.JSON(http.StatusOK, response.Response(responseMap))
}