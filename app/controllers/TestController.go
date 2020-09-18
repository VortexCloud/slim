package controllers

import (
	"github.com/gin-gonic/gin"
	"slim/app/services"
)

var TestCtr TestController

type TestController struct {
	BaseController
}

func (tCtr TestController) Test(ctx *gin.Context) {
	testService := &services.TestService{}
	result := testService.GetValue()
	tCtr.SuccessResponse(ctx, result)
}

func (tCtr TestController) UserList(ctx *gin.Context) {
	testService := &services.TestService{}
	result := testService.GetUserList()
	tCtr.SuccessResponse(ctx, result)
}
