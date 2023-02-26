package routers

import "github.com/gin-gonic/gin"

func IndexRouters(r *gin.Engine) {
	indexRouters := r.Group("/")
	{
		indexRouters.GET("/index", func(c *gin.Context) {
			c.String(200, "c index")
		})

		indexRouters.GET("/login", func(c *gin.Context) {
			c.String(200, "c login")
		})
		indexRouters.GET("/reg", func(c *gin.Context) {
			c.String(200, "c register")
		})
	}
}
