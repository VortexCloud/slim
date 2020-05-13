package main

import (
	"github.com/gin-gonic/gin"
	"slim/routers"
)

func main() {
	var err error
	r := gin.Default()
	routers.WebRouter(r)
	err = r.Run()
	if err != nil {
		panic(err)
	}
}
