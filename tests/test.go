package main

import (
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
	route.Run(":8080")
}
