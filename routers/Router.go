package routers

import (
	"github.com/gin-gonic/gin"
	"slim/app/controllers"
)

func SetRouter(engine *gin.Engine) {

	engine.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})

	engine.GET("/hello", controllers.HelloController{}.Hello)

	api := engine.Group("api")
	{
		v1 := api.Group("v1")
		{
			//user := v1.Group("user",gin.BasicAuth(gin.Accounts{}))
			user := v1.Group("user")
			{
				user.GET("/index", controllers.UserCtr.UserList)
				user.POST("/create", controllers.UserCtr.UserCreate)
			}

			v1.GET("/download_file", controllers.FileCtr.DownLoadFile)
			v1.GET("/uploaded_file", controllers.FileCtr.UploadedFile)
		}
	}

}
