package controllers

import (
	"github.com/gin-gonic/gin"
	"slim/app/services"
)

var UserCtr UserController

type UserController struct {
	BaseController
}

func (userCtr UserController) UserList(ctx *gin.Context) {
	result := services.TestService{}.GetUserList()
	userCtr.SuccessResponse(ctx, result)
}
