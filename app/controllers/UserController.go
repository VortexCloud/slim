package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"slim/app/services"
)

var (
	UserCtr UserController
)

type UserController struct {
	BaseController
}

func (userCtr UserController) UserCreate(ctx *gin.Context) {
	name := ctx.DefaultPostForm("name", "")
	email := ctx.DefaultPostForm("email", "")
	password := ctx.DefaultPostForm("password", "")
	//passwordMd5 := md5.Sum([]byte(password))
	userService := services.UserService{}
	result := userService.CreateUser(name, email, password)
	fmt.Printf("email: %s; password: %s;", email, password)
	userCtr.SuccessResponse(ctx, result)
}

func (userCtr UserController) UserList(ctx *gin.Context) {
	result := services.UserService{}.GetUserList()
	userCtr.SuccessResponse(ctx, result)
}
