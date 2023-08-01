package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"log"
)

func main() {

	route := gin.Default()
	route.LoadHTMLFiles("index.html")
	route.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	route.POST("/upload", func(c *gin.Context) {
		// 从请求中读取文件
		file, err := c.FormFile("f1") // 从请求中获取携带的参数一样的
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message" : err.Error(),
			})
		} else{
			// 将读取的文件保存到服务端本地
			// dst := fmt.Sprintf("./%s", file.Filename)
			log.Println(file.Filename)
			dst := path.Join("./StudyFiles", file.Filename)
			err = c.SaveUploadedFile(file, dst)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message" : err.Error(),
				})
				return
			}else {
				c.JSON(http.StatusOK, gin.H{
					"status" : "ok",
				})
			}
			
		}
	})


	route.Run(":8080")
}