package controllers

import (
	"github.com/gin-gonic/gin"
)

type DemoController struct {
	BaseController
}

var (
	DemoCtr DemoController
)

func (demoCtr DemoController) DingTalkAlarmTest(ctx *gin.Context) {
	demoCtr.SuccessResponse(ctx, "")
}
