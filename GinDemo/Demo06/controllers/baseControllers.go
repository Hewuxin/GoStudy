package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"index": "首页",
	})
}
func MuUpload(c *gin.Context) {
	username := c.PostForm("username")
	form, err := c.MultipartForm()
	if err == nil {
		files := form.File["photo"]
		for _, file := range files {
			c.SaveUploadedFile(file, path.Join("./", file.Filename))
			c.JSON(http.StatusOK, gin.H{
				"username": username,
				"filename": file.Filename,
			})
			fmt.Println("下载成功")
		}
	} else {
		fmt.Println(err.Error())
		return
	}
}

func DmUpload(c *gin.Context) {
	username := c.PostForm("username")
	file1, err1 := c.FormFile("photo1")
	file2, err2 := c.FormFile("photo2")
	fmt.Println(err1, err2)
	if err1 == nil {
		c.SaveUploadedFile(file1, path.Join("./", file1.Filename))
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"filename": file1.Filename,
		})
		fmt.Println("下载成功")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error": "上传失败",
		})
		fmt.Println(err2.Error())
	}
	if err2 == nil {
		c.SaveUploadedFile(file2, path.Join("./", file2.Filename))
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"filename": file2.Filename,
		})
		fmt.Println("下载成功")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error": "上传失败",
		})
		fmt.Println(err2.Error())
	}
}
