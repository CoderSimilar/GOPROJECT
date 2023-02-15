package main
//生产者消费者模型
import(
	
)
func CalSquare() {
    src := make(chan int)     //无缓冲队列
    dest := make(chan int, 3) //有缓冲队列
    
    //A协程用于生产
    go func() {
        defer close(src)
        for i := 0; i < 10; i++ {
            src <- i//将i传入无缓冲队列中
        }
    }()
    
    //B协程用于消费
    go func() {
        defer close(dest)
        for i := range src {
            //有缓冲的通道可以解决生产者消费者速度不匹配的问题。
            dest <- i*i
        }
    }()
    
    //最终主协程遍历有缓冲通道，输出0~9的平方。
    for i := range dest {
        //复杂操作
        println(i)
    }
}

func main() {
	CalSquare()
}