package main

// lesson18 中间件

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// HandlerFunc
func indexHandler(c *gin.Context) {
	fmt.Println("index")
	name, ok := c.Get("name")  // 从上下文中取值（跨中间件存取值)
	if !ok{
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

// 定义一个中间件m1:统计请求处理函数的耗时
func m1(c *gin.Context){
	fmt.Println("m1 in...")
	// 计时
	start := time.Now()
	//go funcXX(c.Copy())  // 在funcXX中只能使用c的拷贝
	c.Next()  // 调用后续的处理函数
	//c.Abort() // 阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out...")

}

func m2(c *gin.Context){
	fmt.Println("m2 in...")
	c.Set("name", "q1mi")  // 在上下文c中设置值
	//c.Abort()  // 阻止调用后续的处理函数
	//return
	fmt.Println("m2 out...")
}


func authMiddleware(doCheck bool)gin.HandlerFunc{
	// 连接数据库
	// 或这一些其他准备工作
	return func(c *gin.Context) {
		if doCheck{
			// 存放具体的逻辑
			// 是否登录的判断
			// if 是登录用户
			c.Next()
			// else
			// c.Abort()
		}else{
			c.Next()
		}
	}
}


func main(){
	//r := gin.Default()  // 默认使用了Logger()和Recovery()
	r := gin.New()
	r.Use(m1, m2, authMiddleware(true))  // 全局注册中间件函数m1,m2

	// GET(relativePath string, handlers ...HandlerFunc) IRoutes
	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})

    // 路由组注册中间件方法1：
	//xxGroup := r.Group("/xx", authMiddleware(true))
	//{
	//	xxGroup.GET("/index", func(c *gin.Context) {
	//		c.JSON(http.StatusOK, gin.H{"msg": "xxGroup"})
	//	})
	//}
	//// 路由组注册中间件方法2：
	//xx2Group := r.Group("/xx2")
	//xx2Group.Use(authMiddleware(true))
	//{
	//	xxGroup.GET("/index", func(c *gin.Context) {
	//		c.JSON(http.StatusOK, gin.H{"msg": "xx2Group"})
	//	})
	//}

	r.Run()  // :8080
}


