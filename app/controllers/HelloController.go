package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var HelloCtr HelloController

type HelloController struct{}

func (hCtr HelloController) Hello(ctx *gin.Context) {
	data := map[string]interface{}{
		"code":    0,
		"message": "Hello GoLang !",
		"data":    "",
	}
	ctx.JSON(http.StatusOK, data)
}
