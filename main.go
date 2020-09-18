package main

import (
	"github.com/gin-gonic/gin"
	"slim/app/libs"
	"slim/routers"
)

func main() {
	var err error

	libs.ConnectMySQL()
	r := gin.Default()
	routers.SetRouter(r)
	err = r.Run()
	if err != nil {
		panic(err)
	}
}
