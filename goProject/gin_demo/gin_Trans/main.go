package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	route.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b"
		route.HandleContext(c)
	})

	route.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"massage" : "b",
		})
	})

	route.Run(":8080")
}