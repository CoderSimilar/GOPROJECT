package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"fmt"
)

func m1(c *gin.Context){
	fmt.Println("I am m1, coming in...")
	Now := time.Now()
	// c.Next()会令程序执行下一个函数，执行完之后再返回本函数继续向下执行
	c.Next()
	fmt.Println(time.Since(Now))
	fmt.Println("I am m1, coming out...")

}

func m2(c *gin.Context){
	fmt.Println("I am m2, coming in...")
	c.Set("name", "similar") // 通过中间件获取信息并传给响应
	// c.Abort()会令程序不执行下一个函数，直接继续执行本函数
	// c.Abort()
	// return
	fmt.Println("I am m2, coming out...")
}

func myFunc(c *gin.Context){
	fmt.Println("I am myFunc")
	name, err := c.Get("name") // 获取中间件传过来的值
	//go funcxx(c.Copy()) // 在funcXX中需要传递c的拷贝c.Copy()，否则不能保证数据的并发安全
	if !err {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : name,
	})

}

// 检查用户是否登录
func islogin() bool{
	return true
}

// 工作中常常这样写中间件
// 检查用户是否登录的中间件，加参数docheck表示是否要检查
func authMiddleware(docheck bool)gin.HandlerFunc {
	// 链接数据库
	// 一些其他的准备工作
	return func(c *gin.Context) {
		if docheck {
			// 检查用户是否登录
			if islogin() {
				c.Next()
			}else{
				c.Abort()
			}
		}else {
			c.Next()
		}
	}
}

func main() {

	route := gin.Default()

	// 注册全局中间件
	route.Use(m1, m2, authMiddleware(true))

	// m1就相当于一个中间件
	route.GET("/index", myFunc)

	// // 为路由组注册中间件 方法1
	// rGroup := route.Group("/similar")
	// {
	// 	rGroup.Use(authMiddleware(true))
	// 	rGroup.GET("/index", func(c *gin.Context) {
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"message" : "GET",
	// 		})
	// 	})
	// 	rGroup.POST("/index", func(c *gin.Context) {
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"message" : "POST",
	// 		})
	// 	})
	// }

	// // 为路由组注册中间件 方法2
	// rGroup = route.Group("/similar", authMiddleware(true))
	// {
	// 	rGroup.GET("/index", func(c *gin.Context) {
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"message" : "GET",
	// 		})
	// 	})
	// 	rGroup.POST("/index", func(c *gin.Context) {
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"message" : "POST",
	// 		})
	// 	})
	// }

	route.Run()


}