package controllers

import (
	"github.com/gin-gonic/gin"
	"slim/app/tools"
)

type DemoController struct {
	BaseController
}

var (
	DemoCtr DemoController
)

func (demoCtr DemoController) DingTalkAlarmTest(ctx *gin.Context) {
	tools.DingTalkAlarm()
	demoCtr.SuccessResponse(ctx, "")
}
