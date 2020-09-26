package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloController struct{}

func (hCtr HelloController) Hello(ctx *gin.Context) {
	data := map[string]interface{}{
		"code":    0,
		"message": "Hello GoLang !",
		"data":    nil,
	}
	ctx.JSON(http.StatusOK, data)
}
