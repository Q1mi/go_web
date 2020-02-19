package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
// gorm demo03

// 1. 定义模型
type User struct {
	ID int64
	Name sql.NullString `gorm:"default:'小王子'"`
	Age int64
}

func main() {
	// 连接MySQL数据库
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 2. 把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	// 3. 创建
	u := User{Name: sql.NullString{String:"",Valid:true}, Age:98}  // 在代码层面创建一个User对象
	fmt.Println(db.NewRecord(&u))  // 判断主键是否为空 true
	db.Debug().Create(&u)  // 在数据库中创建了一条q1mi 18的记录
	fmt.Println(db.NewRecord(&u))  // 判断主键是否为空 false
}
