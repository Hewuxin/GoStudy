package controllers

import "github.com/gin-gonic/gin"

func ApiIndex(c *gin.Context) {
	c.String(200, "chouli api index")
	c.String(200, "controllers")
}

func ApiArticle(c *gin.Context) {
	c.String(200, "chouli api'article")
	c.String(200, "controllers")
}

func ApiEdit(c *gin.Context) {
	c.String(200, "chouli api'edit")
	c.String(200, "controllers")
}
