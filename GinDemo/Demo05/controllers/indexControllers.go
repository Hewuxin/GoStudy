package controllers

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.String(200, "c index")
}

func Login(c *gin.Context) {
	c.String(200, "c login")
}

func Register(c *gin.Context) {
	c.String(200, "c register")
}
