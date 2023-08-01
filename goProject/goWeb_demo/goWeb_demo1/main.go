package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func sayHello(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "Hello, World!",
// 	})
// }

func main() {
	r := gin.Default() // 返回默认的路由引擎

	// 指定用户使用GET请求访问/hello时，执行sayHello这个函数
	r.GET("/book", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"mathod": "GET",
		})
	});

	r.POST("/book", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"mathod": "POST",
		})
	})

	r.PUT("/book", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"mathod": "PUT",
		})
	})

	r.DELETE("/book", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"mathod": "DELETE",
		})
	})

	//启动服务
	r.Run(":9090") // 可以通过指定端口访问

}