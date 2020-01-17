package controllers

import (
	"cocoyo/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type NodeController struct {}

func (c *NodeController) Index(ctx *gin.Context) {
	all, _ := strconv.ParseBool(ctx.DefaultQuery("all", "0"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	nodes := models.GetNodes(all, page, limit)

	ctx.JSON(200, gin.H{"data" : nodes})
}
