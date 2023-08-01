package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func function(w http.ResponseWriter, r *http.Request) {
	// 定义一个函数kua
	// 要么只有一个返回值，要么有两个返回值，第二个返回值必须是err类型
	kua := func(name string)(string, error) {
		return name + "唱，跳, rap, 篮球", nil
	}
	// 定义模板
	t := template.New("f.tml")
	// 在解析模板之前告诉模板引擎我现在多了一个自定义的函数，名字叫kua
	t.Funcs(template.FuncMap {
		"kua99" : kua,
	})
	

	// 解析模板
	_, err :=t.ParseFiles("./f.tml")
	if err != nil {
		fmt.Printf("parse template failed, err : %v\n", err)
	}

	// 渲染模板
	name := "kunkun"
	t.Execute(w, name)
}

func demo1(w http.ResponseWriter, r *http.Request) {
	// 定义模板

	
	// 解析模板
	// 当解析的模板包含嵌套模板时，把被包含的模板卸载包含模板的后面
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err : %v\n", err)
		return
	}

	// 渲染模板
	name := "坤坤"
	t.Execute(w, name);

}

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./index.tml")
	if err != nil {
		fmt.Printf("Parse template failed, err : %v\n", err)
		return
	}

	//渲染模板
	msg := "kunkun"

	t.Execute(w, msg)
}

func home(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./home.tml")
	if err != nil {
		fmt.Printf("Parse template failed, err : %v\n", err)
		return
	}

	//渲染模板
	msg := "kunkun"

	t.Execute(w, msg)
}

func index2(w http.ResponseWriter, r *http.Request) {
	// 定义模板

	// 解析模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index_block.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err : %v\n", err)
		return
	}

	// 渲染模板

	name := "kunkun"
	t.ExecuteTemplate(w, "index_block.tmpl", name)

}

func home2(w http.ResponseWriter, r *http.Request) {
	// 定义模板

	// 解析模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home_block.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err : %v\n", err)
		return
	}

	// 渲染模板

	name := "kunkun"
	t.ExecuteTemplate(w, "home_block.tmpl", name)

}

func main() {
	http.HandleFunc("/", function)
	http.HandleFunc("/Demo", demo1)
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)

	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		fmt.Printf("HTTP server start failed, err : %v\n", err)
		return
	}

}