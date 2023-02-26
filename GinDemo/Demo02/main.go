package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Age      string `form:"age" json:"age"`
}

func main() {
	r := gin.Default()

	r.GET("/json1", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"map": "string",
		})
	})
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "json",
			"msg":     "hahahah ",
		})
	})
	r.GET("/json2", func(c *gin.Context) {
		article := &Article{
			Content: "contents",
			Desc:    "desc",
			Title:   "this is article",
		}
		c.JSON(200, article)
	})
	// 结构体输出
	r.GET("/json3", func(c *gin.Context) {
		article := &Article{
			Content: "post contents",
			Desc:    "p desc",
			Title:   "p this is article",
		}
		c.JSON(200, article)
	})
	// 响应JSONP请求
	//localhost:8080/jsonp?callback=xxx
	//xxx({
	//	"title": "p this is article",
	//	"desc": "p desc",
	//	"content": "jsonp post contents"
	//})
	r.GET("/jsonp", func(c *gin.Context) {
		article := &Article{
			Content: "jsonp post contents",
			Desc:    "p desc",
			Title:   "p this is article",
		}
		c.JSONP(200, article)
	})

	// Get 请求传值
	r.GET("/", func(c *gin.Context) {
		// Query接受get传递的参数
		username := c.Query("username")
		password := c.DefaultQuery("password", "000")
		c.String(200, "值: %v", "首页")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	r.GET("/article", func(c *gin.Context) {
		id := c.DefaultQuery("id", "1")
		c.JSON(http.StatusOK, gin.H{
			"article id": id,
		})
	})

	// POST传值
	r.POST("/user", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "20")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	// 将get或post传过来的值绑定到结构体
	r.GET("/getUser", func(c *gin.Context) {
		// Query接受get传递的参数
		var userInfo UserInfo
		if err := c.ShouldBind(&userInfo); err == nil {
			c.JSON(http.StatusOK, userInfo)
		} else {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		}
		username := c.Query("username")
		password := c.DefaultQuery("password", "000")
		c.String(200, "值: %v", "首页")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	fmt.Println(r.Run())
}
