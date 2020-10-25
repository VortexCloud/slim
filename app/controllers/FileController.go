package controllers

import (
	"github.com/gin-gonic/gin"
)

var (
	FileCtr FileController
)

type FileController struct {
	BaseController
}

func (fileCtr FileController) DownLoadFile(ctx *gin.Context) {

}
