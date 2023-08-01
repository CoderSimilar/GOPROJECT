package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

type UserInfo struct {
	Name string		`form:"username"`
	Password string	`form:"password"`
}
type UserInfo2 struct {
	Name string		`json:"name`
	Identity string	`json:"identity"`
	Hobby string	`json:"hobby"`
}

func main() {

	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.GET("/", func(c *gin.Context) {
		// username := c.Query("username")
		// password := c.Query("password")
		// u := UserInfo {
		// 	Name: username,
		// 	Password: password,
		// }
		var u UserInfo
		err := c.ShouldBind(&u)	// 绑定，ShouldBind把请求里面相关的值取出来直接传递给变量u
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
		}else {
			fmt.Printf("%#v\n", u)

			c.JSON(http.StatusOK, gin.H{
				"status" : "ok",
			})
		}
		
	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/form", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u)	// 绑定，ShouldBind把请求里面相关的值取出来直接传递给变量u
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
		}else {
			fmt.Printf("%#v\n", u)

			c.JSON(http.StatusOK, gin.H{
				"status" : "ok",
			})
		}
	})

	r.POST("/json", func(c *gin.Context) {
		var u UserInfo2
		err := c.ShouldBind(&u)	// 绑定，ShouldBind把请求里面相关的值取出来直接传递给变量u
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
		}else {
			fmt.Printf("%#v\n", u)

			c.JSON(http.StatusOK, gin.H{
				"status" : "ok",
			})
		}
	})

	r.Run(":8080")

}