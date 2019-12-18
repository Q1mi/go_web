package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Gender string
	Age int
}

func sayHello(w http.ResponseWriter, r *http.Request){
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("parse template failed,err:%v", err)
		return
	}
	// 渲染模板
	u1 := User{  // u1.Name
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	m1 := map[string]interface{}{
		"name": "小王子",
		"gender": "男",
		"age": 18,
	}
	hobbyList := []string{
		"篮球",
		"足球",
		"双色球",
	}
	t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
		"hobby": hobbyList,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err:%v", err)
		return
	}
}
