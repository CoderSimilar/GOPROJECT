package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// 路由表示处理某条路径上的请求的一段逻辑方法
	route := gin.Default()

	route.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message" : "GET",
		})
	})

	route.PUT("/index", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message" : "PUT",
		})
	})

	// 使用Any处理任意请求
	route.Any("/any", func(c *gin.Context){
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{
				"Method" : "GET",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"Method" : "POST",
			})
		case http.MethodConnect:
			c.JSON(http.StatusOK, gin.H{
				"Method" : "CONNECT",
			})
		case http.MethodDelete:
			c.JSON(http.StatusOK, gin.H{
				"Method" : "DELETE",
			})
		case http.MethodHead:
			c.JSON(http.StatusOK, gin.H{
				"Method" : "HEAD",
			})
		default:
			c.JSON(http.StatusNotFound, gin.H{
				"Method" : "NOT FOUND",
			})

		}
	})

	// 路由组，可以访问前缀路由下面的所有路由
	rGroup := route.Group("/similar")
	{	
		rGroup.GET("/index", func(c *gin.Context){
			c.JSON(http.StatusOK, gin.H{
				"message" : "GET",
			})
		})
		rGroup.PUT("/index", func(c *gin.Context){
			c.JSON(http.StatusOK, gin.H{
				"message" : "PUT",
			})
		})
		rGroup.POST("/index", func(c *gin.Context){
			c.JSON(http.StatusOK, gin.H{
				"message" : "POST",
			})
		})
	}

	route.Run(":8080")
}