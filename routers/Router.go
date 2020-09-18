package routers

import (
	"github.com/gin-gonic/gin"
	"slim/app/controllers"
)

func SetRouter(engine *gin.Engine) {
	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	engine.GET("/test", controllers.TestCtr.Test)

	v1 := engine.Group("v1")
	{
		user := v1.Group("user")
		{
			user.GET("/index", controllers.TestCtr.UserList)
		}
	}
}
