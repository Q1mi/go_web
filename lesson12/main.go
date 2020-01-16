package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
// 获取form表单提交的参数


func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	// /login post
	r.POST("/login", func(c *gin.Context) {
		// 获取form表单提交的数据
		//username := c.PostForm("username")
		//password := c.PostForm("password")  // 取到就返回值，取不到返回空字符串
		//username := c.DefaultPostForm("username", "somebody")
		//password := c.DefaultPostForm("xxx", "***")
		username, ok := c.GetPostForm("username")
		if !ok {
			username = "sb"
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "***"
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name": username,
			"Password": password,
		})
	})

	r.Run(":9090")
}
