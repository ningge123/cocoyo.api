package controllers

import (
	"cocoyo/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
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