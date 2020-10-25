package main

import (
	"github.com/gin-gonic/gin"
	"slim/routers"
)

func main() {
	var err error

	gin.SetMode(gin.DebugMode)

	//mysql := libs.ConnectMySQL()
	r := gin.Default() //默认使用 Logger 和 Recovery 中间件
	routers.SetRouter(r)
	err = r.Run()
	if err != nil {
		panic(err)
	}
}
