package controllers

import (
	"github.com/gin-gonic/gin"
	"slim/app/services"
)

var TestCtr TestController

type TestController struct {
	BaseController
}

func (tCtr TestController) UserList(ctx *gin.Context) {
	result := services.TestService{}.GetUserList()
	tCtr.SuccessResponse(ctx, result)
}
