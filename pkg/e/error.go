package e

import (
	"cocoyo/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler() gin.HandlerFunc  {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				New(ERROR, fmt.Sprintf("%s", err))
				context.JSON(http.StatusInternalServerError, response.ServerError())
			}
		}()

		context.Next()
	}
}