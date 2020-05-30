package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
	ResponseData
}

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (base *BaseController) SuccessResponse(ctx *gin.Context, data interface{}) {
	base.SetResponse(ctx, 0, "success", data)
}

func (base *BaseController) ErrorResponse(ctx *gin.Context, err error) {
	//log.Print(err)
	base.SetResponse(ctx, 1, "error", err)
}

func (base *BaseController) SetResponse(ctx *gin.Context, code int, message string, data interface{}) {
	base.SetCode(code)
	base.SetMessage(message)
	base.SetData(data)
	ctx.JSON(http.StatusOK, base.ResponseData)
	return
}

func (base *BaseController) SetCode(code int) {
	base.Code = code
}

func (base *BaseController) SetMessage(message string) {
	base.Message = message
}

func (base *BaseController) SetData(data interface{}) {
	base.Data = data
}
