package controllers

import "github.com/gin-gonic/gin"

func getMiddleValue(c *gin.Context) {
	url, _ := c.Get("url")
	//stringUrl, ok := url.(string)
	//if ok {
	//	c.JSON(200, gin.H{
	//		"url": stringUrl,
	//	})
	//} else {
	//	c.String(404, "token获取失败")
	//}
	c.JSON(200, gin.H{
		"url": url,
	})
}

func AdminIndex(c *gin.Context) {

	c.String(200, "chouli admin index")
	c.String(200, "controllers")
	getMiddleValue(c)
}

func AdminUser(c *gin.Context) {
	c.String(200, "chouli admin'user")
	c.String(200, "controllers")
	getMiddleValue(c)

}

func AdminUpdate(c *gin.Context) {
	c.String(200, "chouli admin'update")
	c.String(200, "controllers")
	getMiddleValue(c)

}
