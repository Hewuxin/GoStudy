package routers

import "github.com/gin-gonic/gin"

func ApiRouters(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.String(200, "c api'index")
		})
		apiRouters.GET("/article", func(c *gin.Context) {
			c.String(200, "c api'article")
		})
		apiRouters.GET("/edit", func(c *gin.Context) {
			c.String(200, "c api edit")

		})
	}
}
