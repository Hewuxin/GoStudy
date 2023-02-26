package routers

import "github.com/gin-gonic/gin"

func AdminRouters(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", func(c *gin.Context) {
			c.String(200, "chouli admin index")
		})
		adminRouters.GET("/user", func(c *gin.Context) {
			c.String(200, "chouli admin'user")
		})
		adminRouters.GET("/update", func(c *gin.Context) {
			c.String(200, "chouli admin'update")
		})
	}
}
