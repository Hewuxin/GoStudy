package main

import (
	"Demo06/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 单个/多个文件 文件上传
// 设置cookie

func main() {
	router := gin.Default()

	router.GET("/", controllers.Index)
	router.POST("/upload", controllers.DmUpload)
	router.POST("/muUpload", controllers.MuUpload)

	router.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "Not set cookie"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	})
	router.Run()
}
