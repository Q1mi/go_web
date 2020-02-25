package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
// gorm demo06

// 1. 定义模型
type User struct {
	//gorm.Model  // ID CreateAt UpdateAt DeleteAt
	ID int64
	Name string
	Age int64
	Active bool
}

func main() {
	// 2.连接MySQL数据库
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 3. 把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})
	// 4. 创建
	//u1 := User{Name:"q1mi2", Age:18, Active:true}
	//db.Create(&u1)
	//u2 := User{Name:"jinzu2", Age:20, Active:false}
	//db.Create(&u2)

	// 5.删除
	//var u = User{}
	////u.ID = 1
	//u.Name = "jinzu2"
	//db.Debug().Delete(&u)

	//db.Debug().Where("name=?", "jinzu2").Delete(User{})
	//db.Delete(User{}, "age=?", 18)

	var u1 User
	db.Debug().Where("name=?", "jinzu2").Find(&u1)
	fmt.Println(u1)

	// 物理删除
	//db.Debug().Unscoped().Where("name=?", "q1mi2").Delete(User{})
}
