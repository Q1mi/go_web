package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// lesson17 路由与路由组

func main() {
	r := gin.Default()
	// 访问/index的GET请求会走这一条处理逻辑
	// 路由
	//r.HEAD()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "delete",
		})
	})
	// Any：请求方法的大集合/大杂烩
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{"method": "GET"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "POST"})
			// ...
		}
	})
	// NoRoute
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "liwenzhou.com"})

	})

	// 视频的首页和详情页
	r.GET("/video/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg":"/video/index"})
	})
	r.GET("/video/xx", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg":"/video/xx"})
	})
	r.GET("/video/oo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg":"/video/oo"})
	})
	// 路由组的组 多用于区分不同的业务线或API版本
	// 把公用的前缀提取出来，创建一个路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg":"/video/index"})
		})
		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg":"/video/xx"})
		})
		videoGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg":"/video/oo"})
		})
	}

	// 商城的首页和详情页
	r.GET("/shop/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg":"/shop/index"})
	})
	r.GET("/shop/xx", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg":"/shop/xx"})
	})
	r.GET("/shop/oo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg":"/shop/oo"})
	})

	r.Run(":9090")
}
