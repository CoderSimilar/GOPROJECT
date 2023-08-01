package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	// querystring

	// GET请求 URL ?后面的是querystring参数
	// key=value格式，多个key-value使用&连接
	// eq: /web/query=蔡徐坤&age=18
	route.GET("/web", func(c *gin.Context) {

		// 获取浏览器那边发送的请求携带的querystring参数
		// name := c.Query("query")	// 通过Query获取请求中的querystring参数
		// name := c.DefaultQuery("query", "somebody")	// 当查询不到时采用默认值
		
		name, ok := c.GetQuery("query")	// 取到返回(值，true)，取不到返回("", false)
		if !ok {
			// 如果取不到值，设置默认值并返回
			name = "somebody"
		}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})

	route.Run(":9090")
}