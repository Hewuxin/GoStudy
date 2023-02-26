package main

import (
	"Demo04/routers"
	"github.com/gin-gonic/gin"
)

// Gin 路由分组 和 路由抽离
func main() {
	router := gin.Default()
	// 1.路由分组
	//indexRouters := router.Group("/")
	//{
	//	indexRouters.GET("/index", func(c *gin.Context) {
	//		c.String(200, "index")
	//	})
	//
	//	indexRouters.GET("/login", func(c *gin.Context) {
	//		c.String(200, "login")
	//	})
	//	indexRouters.GET("/reg", func(c *gin.Context) {
	//		c.String(200, "register")
	//	})
	//}
	//adminRouters := router.Group("/admin")
	//{
	//	adminRouters.GET("/", func(c *gin.Context) {
	//		c.String(200, "admin index")
	//	})
	//	adminRouters.GET("/user", func(c *gin.Context) {
	//		c.String(200, "admin'user")
	//	})
	//	adminRouters.GET("/update", func(c *gin.Context) {
	//		c.String(200, "admin'update")
	//	})
	//}
	//apiRouter := router.Group("/api")
	//{
	//	apiRouter.GET("/", func(c *gin.Context) {
	//		c.String(200, "api'index")
	//	})
	//	apiRouter.GET("/article", func(c *gin.Context) {
	//		c.String(200, "api'article")
	//	})
	//	apiRouter.GET("/edit", func(c *gin.Context) {
	//		c.String(200, "api edit")
	//
	//	})
	//}

	// 2.路由抽离
	routers.AdminRouters(router)
	routers.ApiRouters(router)
	routers.IndexRouters(router)
	router.Run()
}
