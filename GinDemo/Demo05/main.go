package main

// 中间件
import (
	"Demo05/routers"
	"github.com/gin-gonic/gin"
)

func GlobalMiddleWareOne(c *gin.Context) {
	c.String(200, "Before  ")
	c.Next()
	c.String(200, "  After")
}

func GlobalMiddleWareTwo(c *gin.Context) {
	c.String(200, "TwoBefore  ")
	c.Next()
	c.String(200, "  TwoAfter")
}

// Gin 控制器 中间件
func main() {
	router := gin.Default()
	// 使用全局中间件
	//router.Use(GlobalMiddleWareOne, GlobalMiddleWareTwo)
	// 2.路由抽离
	routers.AdminRouters(router)
	routers.ApiRouters(router)
	routers.IndexRouters(router)
	router.Run()
}
