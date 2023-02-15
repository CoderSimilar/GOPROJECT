package main

import (
	"bufio"
	"log"
	"net"
)
//TCP echo server 简单实现，给他发送什么，他就给你回复什么
func main() {
	server, err := net.Listen("tcp", "127.0.0.1:1080")//监听端口，返回一个server

	if err != nil {
		panic(err)
	}

	for {
		client, err := server.Accept()//接受一个请求，成功的话会返回一个连接
		if err != nil {
			log.Printf("Accept failed %v", err);
			continue;
		}
		go process(client)//使用goroutine启用协程处理这个连接
	}
}

func process(conn net.Conn) {
	// conn, err := net.Dial("tcp", "127.0.0.1:1080")//和127.0.0.1建立连接
	// if err != nil {
	// 	log.Fatalf("Failed to connect to server : %v", err)
	// }
	defer conn.Close()//在函数调用结束的时候关闭连接
	reader := bufio.NewReader(conn)//基于此链接创建一个制度的带缓冲的流

	for {
		b, err := reader.ReadByte()//每次读取一个字节
		if err != nil {
			break
		}
		_, err = conn.Write([]byte{b})//将这个字节写入
		if err != nil {
			break
		}
	}

}