package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Event struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {

	// 创建数据库
	// CREATE DATABASE SimTodoList

	// 连接数据库
	dsn := "similar:123456@tcp(127.0.0.1:3306)/SimTodoList?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 数据库和结构体相对应
	db.AutoMigrate(&Event{})

	// 注册路由
	route := gin.Default()
	// 加载静态文件
	route.Static("/static", "./dist/static")
	// 加载HTML文件
	route.LoadHTMLGlob("templates/*")

	route.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)

	})

	// 代办事项

	route_v1 := route.Group("/v1")
	{
		// 添加
		route_v1.POST("/todo", func(c *gin.Context) {
			// 前端页面填写待办事项，点击提交，会发起请求到这里
			// 1，获取前端传来的JSON数据
			var event Event
			err = c.ShouldBind(&event)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%#v\n", event)
			// 创建事件并添加到数据库中
			// 返回响应
			if err = db.Create(&event).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, event)
				// 实际上一般这么写：
				// c.JSON(http.StatusOK, gin.H{
				// 	"message": "添加成功",
				// 	"data" : event,
				// 	"code" : 2000,
				// 	....
				// })
			}
		})
		// 查看
		route_v1.GET("/todo", func(c *gin.Context) {
			// 查询events这个表里的所有数据
			// 创建一个切片
			var eventList []Event
			if err = db.Find(&eventList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				// fmt.Printf("%#v\n", eventList)
				c.JSON(http.StatusOK, eventList)
			}

		})
		// 修改
		route_v1.PUT("/todo/:id", func(c *gin.Context) {
			// 获取参数
			id, ok := c.Params.Get("id")
			fmt.Println(id)
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "id不能为空",
				})
				return
			}
			var event Event // 定义一个结构体变量
			// 根据传进来的id查询到要修改的变量
			if err = db.Where("id=?", id).First(&event).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.BindJSON(&event)
			if err = db.Save(&event).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, event)
			}

		})
		// 删除
		route_v1.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id") // 获取要删除的id
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "id不能为空",
				})
				return
			}
			if err = db.Where("id=?", id).Delete(&Event{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"id": "deleted",
				})
			}
		})
	}

	// 运行路由
	route.Run(":8080")

}
