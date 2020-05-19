package main

import (
	"github.com/gin-gonic/gin"
	"slim/databases"
	"slim/routers"
)

func main() {
	var err error

	err = databases.MysqlInstance()

	r := gin.Default()
	routers.Router(r)
	err = r.Run()
	if err != nil {
		panic(err)
	}
}
