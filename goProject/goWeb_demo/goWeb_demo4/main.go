package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func xss(w http.ResponseWriter, r *http.Request) {
	// 创建模板
	// 解析模板
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap {
		"safe" : func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("./xss.tmpl")

	if err != nil {
		fmt.Printf("Parse template failed, err : %v\n", err)
		return
	}

	// 渲染模板
	script1 := `<script>alert("你好, 我是kunkun, 喜欢唱、跳、rap、篮球");</script>`
	script2 := `<a href="http://www.baidu.com">百度</a>`
	t.Execute(w, script1)
	t.Execute(w, script2)
}

func main() {
	http.HandleFunc("/xss", xss)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		fmt.Printf("HTTP server start failed, err : %v\n", err)
		return
	}

	
}