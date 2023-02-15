package main

import (
	"fmt"
	"time"
)

//快速打印hello goroutinue :0 ~ hello goroutinue :4
func hello(i int) {

	fmt.Println("hello goroutinue : " + fmt.Sprint(i))
}
func main() {
		for i := 0; i < 5; i++ {
			go func(j int){
				hello(j)
			}(i)
		}
		time.Sleep(time.Second)//使用了暴力的sleep做了阻塞，防止子协程打印完之前主协程不会退出
	}