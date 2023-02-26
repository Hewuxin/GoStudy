package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认的路由引擎
	fmt.Println("This is Gin demo01")
	r := gin.Default()
	// 配置路由
	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "值:%v", "你好 gin\n")
		c.JSON(200, gin.H{
			"message": "po",
		})
	})
	r.POST("/post", func(n *gin.Context) {
		n.JSON(200, gin.H{
			"message": "这是post请求",
		})
	})
	r.PUT("/put", func(n *gin.Context) {
		n.JSON(200, gin.H{
			"message": "这是put请求",
		})
	})
	r.DELETE("/delete", func(n *gin.Context) {
		n.JSON(200, gin.H{
			"message": "这是delete请求",
		})
	})
	// 启动一个web服务
	panic(r.Run(":8000")) // listen and serve on 0.0.0.0:8080
}
