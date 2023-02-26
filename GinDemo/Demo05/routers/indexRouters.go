package routers

import (
	"Demo05/controllers"
	"Demo05/middleWares"
	"github.com/gin-gonic/gin"
)

func IndexRouters(r *gin.Engine) {
	indexRouters := r.Group("/", middleWares.GoroutineTestMiddleWares)
	{
		indexRouters.GET("/", controllers.Index)

		indexRouters.GET("/login", controllers.Login)
		indexRouters.GET("/reg", controllers.Register)
	}
}
