package main

import (
	"github.com/gin-gonic/gin"
	"slim/routers"
)

func main() {
	var err error

	gin.SetMode(gin.DebugMode)

	//mysql := libs.ConnectMySQL()
	r := gin.Default()
	routers.SetRouter(r)
	err = r.Run()
	if err != nil {
		panic(err)
	}
}
