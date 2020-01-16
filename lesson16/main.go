package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"status": "ok",
		//})
		// 跳转到 sogo.com
		c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com")
	})

	r.GET("/a", func(c *gin.Context) {
		// 跳转到 /b 对应的路由处理函数
		c.Request.URL.Path = "/b"  // 把请求的URI修改
		r.HandleContext(c)  // 继续后续的处理
	})
	r.GET("/b", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "b",
			})
	})
	r.Run(":8080")
}
