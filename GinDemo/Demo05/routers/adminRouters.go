package routers

import (
	"Demo05/controllers"
	"Demo05/middleWares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminMiddleWares(C *gin.Context) {
	C.JSON(http.StatusOK, gin.H{
		"title": "middleWares",
	})
	C.Next()
	C.String(200, "/n middleWare ending...")
}

func AdminRouters(r *gin.Engine) {
	//adminRouters := r.Group("/admin", AdminMiddleWares)
	adminRouters := r.Group("/admin", middleWares.TestMiddleWares)
	{
		adminRouters.GET("/", controllers.AdminIndex)
		adminRouters.GET("/user", controllers.AdminUser)
		adminRouters.GET("/update", controllers.AdminUpdate)
	}
}
