package main

import (
	"bufio"
	"encoding/binary"
	"errors"
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
		log.Printf("client %v auth failed:%v", conn.RemoteAddr(), err)
		return
	}
	log.Print("auth secess")

	err = connect(reader, conn)

	if err != nil {
		log.Printf("client %v connect failed:%v", conn.RemoteAddr(), err)
		return
	}
	log.Print("connect success")


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
//试图读取用户浏览器发送的一个报文，报文包含用户需要访问的url或者ip+端口。
func connect(reader *bufio.Reader, conn net.Conn) (err error) {
	// +----+-----+-------+------+----------+----------+
	// |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER 版本号，socks5的值为0x05
	// CMD 0x01表示CONNECT请求,让代理服务器和下游服务器创建连接
	// RSV 保留字段，值为0x00
	// ATYP 目标地址类型，DST.ADDR的数据对应这个字段的类型。
	//   0x01表示IPv4地址，DST.ADDR为4个字节
	//   0x03表示域名，DST.ADDR是一个可变长度的域名，第一个字节是长度，后面的字节是真正的域名
	// DST.ADDR 一个可变长度的值
	// DST.PORT 目标端口，固定2个字节

	//创建一个长度为4的缓冲区，使用io.ReadFull将其填充，就可以一次性读取到前面四个字节的内容。
	buf := make([]byte, 4)
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		return fmt.Errorf("read header failed!%w", err)
	}
	ver, cmd, atyp := buf[0], buf[1], buf[3]
	if ver != socks5Ver {
		return fmt.Errorf("not supposed ver:%v", err)
	}
	if cmd != cmdBind {
		return fmt.Errorf("not supposed cmd:%v", cmd)
	}
	addr := ""
	switch atyp {
	case atypIPV4://如果是IPV4请求,将IP地址填充到buf里并打印
		_, err = io.ReadFull(reader, buf)
		if err != nil {
			return fmt.Errorf("read atyp failed:%w", err)
		}
		addr = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2],buf[3])
	case atypeHOST://如果是HOST请求，先读取HOST请求的长度，然后创建HOST长度的缓冲区，然后将其转换成字符串
		hostsize, err := reader.ReadByte()//先读取第一个字节，即HOST请求的长度
		if err != nil {
			return fmt.Errorf("read hostsize failed:%w", err)
		}
		host := make([]byte, hostsize)//创建host请求长度的切片
		_, err = io.ReadFull(reader, host)//将其填充满
		if err != nil {
			return fmt.Errorf("read host failed:%w", err)
		}
		addr = string(host)//将host请求转换成字符串
	case atypeIPV6://如果是IPV6的话也是读取固定长度
		_, err = io.ReadFull(reader, buf)
		if err != nil {
			return fmt.Errorf("read atyp failed:%w", err)
		}
		addr = fmt.Sprintf("%s.%s.%s.%s", buf[0], buf[1], buf[2], buf[3])
	default://其他方式直接打印错误提示
		return errors.New("invalid atyp")
	}//至此，前面的五个字段数据已经读取完成，只剩下最后一个端口号字段，长度为固定两个字节

	//复用前面的buf缓冲区，使用切片对其裁剪
	//新的切片和原始切片是互用底层数据的，所以在缓冲区内能够直接读到端口号数据
	_, err = io.ReadFull(reader, buf[:2])
	if err != nil {
		return fmt.Errorf("read port failed:%w", err)
	} 
	//使用binary包内的函数采用大端字节的方式读取出整型数字
	port := binary.BigEndian.Uint16(buf[:2])
	//打印日志，表示将会与这个地址的这个端口进行连接
	log.Println("dial", addr, port)
	

	//按照协议，接收到浏览器的请求后，需要进行回包
	// +----+-----+-------+------+----------+----------+
	// |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER socks版本，这里为0x05
	// REP Relay field,内容取值如下 X’00’ succeeded
	// RSV 保留字段
	// ATYPE 地址类型
	// BND.ADDR 服务绑定的地址，并非所支持的connect请求所必需的，可直接填成0值
	// BND.PORT 服务绑定的端口DST.PORT，并非所支持的connect请求所必需的，可直接填成0值

	//第一个代表服务器版本号，5；
	//第二个REP填0，代表成功；
	//第三个RSV保留字段填0；
	//第四个ATYP填1，代表使用IPV4
	//第五个字段ADDR不用，4个字节4个0
	//第六个端口号不需要，2个字节2个0
	_, err = conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0, 0, 0, 0, 0, 0})
	if err != nil {
		return fmt.Errorf("write failed: %w", err)
	}
	return nil
}