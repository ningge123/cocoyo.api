package controllers

import (
	"github.com/gin-gonic/gin"
)

type NodeController struct {}

func (c *NodeController) Index(ctx *gin.Context) {
	//all, _ := strconv.ParseBool(ctx.DefaultQuery("all", "0"))
	//page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	//limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	//
	//nodes, err := models.GetNodes(all, page, limit)
	//
	//if err != nil {
	//	e.New(e.ERROR, err.Error())
	//}
	//
	//ctx.JSON(200, gin.H{"data" : nodes})
}
