package main

import (
	"fmt"
	"net/http"
	"html/template"
)
// 遇事不决 写注释

func sayHello(w http.ResponseWriter, r *http.Request){
	// 2. 解析模板
	t, err := template.ParseFiles("./hello.tmpl")  // 请勿刻舟求剑
	if err != nil {
		fmt.Println("Parse template failed,err:%v", err)
		return
	}
	// 3. 渲染模板
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("render template failed,err:%v", err)
		return
	}

}


func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err:%v", err)
		return
	}
}
