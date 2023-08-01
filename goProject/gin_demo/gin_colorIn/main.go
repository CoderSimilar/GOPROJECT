package main

import (
	"html/template"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	route := gin.Default(); // route是gin的一个默认路由

	//加载静态文件
	// 第一个参数是在HTML文件中填写的静态文件目录，第二个参数是你的静态文件目录
	route.Static("/xxx", "./statics")

	//自定义模板函数
	route.SetFuncMap(template.FuncMap{
		"safe" : func(s string) template.HTML {
			return template.HTML(s)
		},
	})
	
	//gin使用 LoadHTMLGlob 或 LoadHTMLFiles 加载HTML文件
	//LoadHTMLFiles可以通过文件名对多个HTML文件进行加载
	// route.LoadHTMLFiles("./templates/posts/index.html", "./templates/users/index.html")
	//LoadHTMLGlob可以通过正则表达式的方式对多个HTML文件进行加载
	route.LoadHTMLGlob("./templates/**/*.html")
	
	//定义一个GET请求的路由处理函数
	route.GET("/posts/index", func(c *gin.Context) {
		//使用c.HTML来渲染HTML模板并返回给客户端
		// 第二个参数是渲染的文件名，如果在文件中使用{{define NAME}}指定了名字的话，那么就是NAME，如果没有指定名字，那么就是该文件的文件名
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "<a href=\"http://www.baidu.com\">baidu</a>",
			"medium": "I am Iron Man",
		})
		// c.HTML(http.StatusOK, "posts/index.html", gin.H{
		// 	"midum": "I am Iron Man",
		// })
	})

	route.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index.tmpl",
		})
	})

	route.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	route.Run(":9090")

}