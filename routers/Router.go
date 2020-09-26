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

	engine.GET("/hello", controllers.HelloCtr.Hello)

	api := engine.Group("api")
	{
		v1 := api.Group("v1")
		{
			user := v1.Group("user")
			{
				user.GET("/index", controllers.TestCtr.UserList)
			}
		}
	}

}
