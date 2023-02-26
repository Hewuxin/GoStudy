package main

// Gin get post 传值
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type Article struct {
	Title string `json:"title" form:"title"`
	Desc  string `json:"desc" form:"desc"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "index")
	})

	r.GET("/login", func(c *gin.Context) {
		user := &UserInfo{}
		fmt.Printf("%#v", user)
		if err := c.ShouldBind(user); err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		}
	})
	// 将POST接收的值绑定到结构体
	r.POST("/reg", func(c *gin.Context) {
		article := &Article{}
		if err := c.ShouldBind(article); err == nil {
			c.JSON(http.StatusOK, article)
			fmt.Println(article.Title)
			fmt.Println(article.Desc)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		}
	})

	// 使用DefaultPostForm接收POST传值，然后将其赋值给结构体
	r.POST("/test", func(c *gin.Context) {
		article := &Article{}
		title := c.DefaultPostForm("title", "nil")
		desc := c.DefaultPostForm("desc", "nil")
		article.Title = title
		article.Desc = desc
		c.JSON(http.StatusOK, article)
	})

	// 动态路由
	r.GET("/user/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		c.String(http.StatusOK, uid)
	})
	r.Run()
}
