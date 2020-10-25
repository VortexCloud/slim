package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Person struct {
	Name     string    `form:"name" binding:"required"`
	Address  string    `form:"address" binding:"required"`
	Birthday time.Time `form:"birthday"`
}

func startPage(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	//lastname := c.Query("lastname")
	//firstname := c.DefaultQuery("firstname", "Guest")
	//message := c.PostForm("message")  //  同上
	//message := c.DefaultPostForm("message")
	if err := c.ShouldBind(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if person.Name != "xueyuanjun" || person.Address != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})

	//c.String(200, "Success")
}

func main() {
	route := gin.Default()
	// GET 请求
	route.GET("/testing", startPage)
	// POST 请求
	route.POST("/testing", startPage)

	// HTTP 重定向
	route.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	})

	// 路由重定向
	route.GET("/test1", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		route.HandleContext(c)
	})
	route.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	autotls.Run(route, "example.test")
	//route.Run(":8080")

	//server02 := &http.Server{
	//	Addr: ":8081",
	//	Handler: router02(),
	//	ReadHeaderTimeout: 5 * time.Second,
	//	WriteTimeout: 10 * time.Second,
	//}
	//server02.ListenAndServe()
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})

	return e
}
