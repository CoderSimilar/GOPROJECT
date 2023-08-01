package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Name string
	Age uint
	Birthday time.Time
}

func main() {

	// 连接mysql数据库，首先填写数据库的用户名，然后填写地址，接着填写操作的数据库的名字，然后填写编码方式，以及是否要解析时间
	dsn := "similar:123456@tcp(127.0.0.1:3306)/test_1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})	// 创建User表，把结构体和表对应

	user := User{Name: "simi", Age : 18, Birthday : time.Now()}
	db.Create(&user)	// 通过数据指针创建表
	fmt.Printf("%#v\n", user)

	// 可以使用Create创建多条记录
	users := [] *User {
		{Name : "roy", Age : 24, Birthday : time.Now()},
		{Name : "xb", Age : 23, Birthday : time.Now()},
		{Name : "dz", Age : 24, Birthday : time.Now()},
		{Name : "jx", Age : 23, Birthday : time.Now()},
	}
	result := db.Create(&users)
	fmt.Println(result.RowsAffected, "rows has been changed") // 打印更新了几行数据

	// 使用指定的字段创建记录
	db.Select("Name", "Age", "Birthday").Create(&user)
	// 创建记录并忽略要省略的传递字段的值
	user.Birthday = time.Now().Add(10000000000)
	db.Omit("Name", "Age", "Brithday").Create(&user)

	// // 更新指定的行
	// db.Model(&User{}).Where("Name = ?", "roy").Update("Name", "Kunkun") // 把roy换成kunkun
	// // 删除指定的行
	// result = db.Model(&User{}).Where("Name = ?", "kunkun").Delete(&User{}) // 把kunkun都删除掉
	// if result.Error != nil {
	// 	panic(result.Error)
	// }else {
	// 	fmt.Println(result.RowsAffected, "rows has been changed")
	// }
	// // 删除整张表
	// db.Migrator().DropTable(&User{})
}