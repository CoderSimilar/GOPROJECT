package main

import(
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}
 
func main() {
	// 链接mysql数据库，首先填写数据库的用户名，然后填写地址，接着填写操作的数据库的名字，然后填写编码方式，以及是否要解析时间
	dsn := "similar:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!= nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	// 创建User表，把结构体和数据表对应
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	u1 := UserInfo{
		Name: "kunkun",
		ID: 1,
		Gender: "man",
		Hobby: "sing, dance, rap, basketball",
	}
	db.Create(&u1) // 通过数据的指针创建
	fmt.Printf("%#v\n", u1)

	// 查询数据
	var u UserInfo
	db.First(&u)	// 查询表中第一条数据保存到u中
	fmt.Printf("%#v\n", u)

	// 更新表
	db.Model(&u).Update("Hobby", "sex")
	db.Model(&u).Update("Hobby", "sing, dance, rap, basketball")
	// 删除，清空表
	db.Delete(&u)
	
}