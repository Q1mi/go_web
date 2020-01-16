package main

// querystring
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// querystring

	// GET请求 URL ?后面的是querystring参数
	// key=value格式，多个key-value用 & 连接
	// eq:  /web/query=小王子&age=18
	r.GET("/web", func(c *gin.Context) {
		// 获取浏览器那边发请求携带的 query string 参数
		name := c.Query("query")  // 通过Query获取请求中携带的querystring参数
		age := c.Query("age")
		//name := c.DefaultQuery("query", "somebody")  // 取不到就用指定的默认值
		//name, ok := c.GetQuery("query")  // 取到返回(值, true)，取不到返回("", false)
		//if !ok {
		//	// 取不到
		//	name = "somebody"
		//}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age": age,
		})
	})

	r.Run(":9090")
}
