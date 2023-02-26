package routers

import (
	"Demo05/controllers"
	"Demo05/middleWares"
	"github.com/gin-gonic/gin"
)

func ApiRouters(r *gin.Engine) {
	apiRouters := r.Group("/api", middleWares.InitMiddleWares)
	{
		apiRouters.GET("/", controllers.ApiIndex)
		apiRouters.GET("/article", controllers.ApiArticle)
		apiRouters.GET("/edit", controllers.ApiEdit)
	}
}
