package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default() // 获取默认路由
	
	r.GET("/JSON", func(c *gin.Context) {
		// c.JSON(http.StatusOK, map[string] interface{}{
		// 	"msg": "hello",
		// 	"code": 200,
		// 	"data": "kunkun",
		// })

		//gin.H相当于 map[string] interface{}
		// c.JSON(http.StatusOK, gin.H {
		// 		"msg": "hello",
		// 		"code": 100,
		// 		"data": "kunkun",
		// 	})
		// })

		//使用结构体传递JSON
		type msg struct {
			Msg string
			Name string
			Age int
		}
		data := msg{
			Msg: "hello",
			Name: "kunkun",
			Age: 18,
		}
		c.JSON(http.StatusOK, data)
})

	r.Run(":8080")

}