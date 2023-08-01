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

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 1，定义模板
	// 2，解析模板

	t, err := template.ParseFiles("./hello.tmpl")

	if err != nil {
		fmt.Printf("Parse template failed, err : %v\n", err)
		return
	}

	// name := "小王子"
	// 3，渲染模板
	u1 := User {
		Name: "坤坤",
		Gender: "男",
		Age: 18,
	}
	m1 := map[string]interface{}{
		"name" : "坤坤",
		"gender": "男",
		"age": 18,
	}

	hobbyList := []string{"唱", "跳", "rap", "篮球"}

	err = t.Execute(w, map[string]interface{}{
		"u1" : u1,
		"m1" : m1,
		"hobby" : hobbyList,
	})
	if err != nil {
		fmt.Printf("Execute template failed, err : %v\n", err)
		return
		
	}

}

func main() {

	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		fmt.Printf("HTTP server start failde, err : %v\n", err)
		return
	}

}