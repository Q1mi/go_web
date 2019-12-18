package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request){
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}
	msg := "小王子"
	// 渲染模板
	t.Execute(w, msg)
}

func home(w http.ResponseWriter, r *http.Request){
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}
	msg := "小王子"
	// 渲染模板
	t.Execute(w, msg)
}

func index2(w http.ResponseWriter, r *http.Request){
	// 定义模板（模板继承的方式）
	// 解析模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index2.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	// 渲染模板
	name := "小王子"
	t.ExecuteTemplate(w, "index2.tmpl", name)
}
func home2(w http.ResponseWriter, r *http.Request){
	// 定义模板（模板继承的方式）
	// 解析模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home2.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	// 渲染模板
	name := "七米"
	t.ExecuteTemplate(w, "home2.tmpl", name)
}


func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err:%v", err)
		return
	}
}
