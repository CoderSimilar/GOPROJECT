package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

const socks5Ver = 0x05
const cmdBind = 0x01
const atypIPV4 = 0x01
const atypeHOST = 0x03
const atypeIPV6 = 0x04

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

	err := auth(reader, conn)
	if err != nil {
		log.Print("client %v auth failed:%v", conn.RemoteAddr(), err)
		return
	}
	log.Print("auth secess")

}

func auth(reader *bufio.Reader, conn net.Conn) (err error) {
	//认证阶段：1，浏览器给代理服务器发送一个报文，报文有三个字段。
	// +----+----------+----------+
	// |VER | NMETHODS | METHODS  |
	// +----+----------+----------+
	// | 1  |    1     | 1 to 255 |
	// +----+----------+----------+
	// VER: 协议版本，socks5为0x05
	// NMETHODS: 支持认证的方法数量
	// METHODS: 对应NMETHODS，NMETHODS的值为多少，METHODS就有多少个字节。RFC预定义了一些值的含义，内容如下:
	// X’00’ NO AUTHENTICATION REQUIRED
	// X’02’ USERNAME/PASSWORD
	ver, err := reader.ReadByte() //读取版本号
	if err != nil {
		return fmt.Errorf("read ver failed:%v", err)
	}
	//如果版本号不对
	if ver != socks5Ver {
		return fmt.Errorf("not support ver:%v", err)
	}
	methodSize, err := reader.ReadByte() //读取
	if err != nil {
		return fmt.Errorf("read methodSize failed!:%v", err)
	}
	method := make([]byte, methodSize)
	_, err = io.ReadFull(reader, method)
	if err != nil {
		return fmt.Errorf("read method failed!%v", err)
	}
	log.Println("ver", ver, "method", method)//打印版本号和方法
	// +----+--------+
	// |VER | METHOD |
	// +----+--------+
	// | 1  |   1    |
	// +----+--------+
	_, err = conn.Write([]byte{socks5Ver, 0x00})
	if err != nil {
		return fmt.Errorf("writer failed!:%v", err)
	}
	return nil
}