package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "status": "no value"})
}
