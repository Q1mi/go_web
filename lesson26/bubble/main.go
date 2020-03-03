package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

// Todo Model
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func initMySQL()(err error){
	dsn := "root:root1234@tcp(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close()  // 程序退出关闭数据库连接
	// 模型绑定
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	
	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {
			// 前端页面填写待办事项 点击提交 会发请求到这里
			// 1. 从请求中把数据拿出来
			var todo Todo
			c.BindJSON(&todo)
			// 2. 存入数据库
			//err = DB.Create(&todo).Error
			//if err!= nil {
			//}
			// 3. 反回响应
			if err = DB.Create(&todo).Error;err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			}else{
				c.JSON(http.StatusOK, todo)
				//c.JSON(http.StatusOK, gin.H{
				//	"code": 2000,
				//	"msg": "success",
				//	"data": todo,
				//})
			}
		})
		// 查看所有的待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			// 查询todo这个表里的所有数据
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err!= nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			}else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		// 查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {
			
		})
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err!=nil{
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err!= nil{
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			}else{
				c.JSON(http.StatusOK, todo)
			}
		})
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			if err = DB.Where("id=?", id).Delete(Todo{}).Error;err!=nil{
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			}else{
				c.JSON(http.StatusOK, gin.H{id:"deleted"})
			}
		})
	}
	r.Run()
}
