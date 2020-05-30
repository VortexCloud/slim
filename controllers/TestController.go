package controllers

import (
	"github.com/gin-gonic/gin"
)

var TestCtr TestController

type TestController struct {
	BaseController
}

func (tCtr TestController) Test(ctx *gin.Context) {
	tCtr.SuccessResponse(ctx, "")
}
