package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	FileCtr FileController
)

type FileController struct {
	BaseController
}

func (fileCtr FileController) DownLoadFile(ctx *gin.Context) {
	response, err := http.Get("https://apk.dalongyun.com/dalongyun_page/v2.0.1/new_moban/ad.jpg")
	if err != nil || response.StatusCode != http.StatusOK {
		ctx.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="gopher.png"`,
	}

	ctx.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	//fileCtr.SuccessResponse(ctx,"")
}

func (fileCtr FileController) UploadedFile(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	// ctx.SaveUploadedFile(file, dst)
	fileCtr.SuccessResponse(ctx, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
