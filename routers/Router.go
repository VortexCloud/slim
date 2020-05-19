package routers

import (
	"github.com/gin-gonic/gin"
	"slim/controllers"
)

func Router(engine *gin.Engine) {
	engine.GET("/test", controllers.Test)
	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "success",
		})
	})
}
