package main

import (
	"fmt"
	"net"
)

func main()  {
	//指定服务器通信协议，ip地址和端口号，创建一个用于监听的socket
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.listen err", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务器端等待客户端连接。。。")
	//阻塞监听连接客户端请求，accept成功连接，会返回一个用于通信的socket
	connect, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("accept err", err1)
		return
	}
	defer connect.Close()

	fmt.Println("连接成功")

	buf := make([]byte, 4096)
	n, err2 := connect.Read(buf)
	if err2 != nil {
		fmt.Println("read err", err2)
		return
	}
	//处理数据，打印
	fmt.Println("服务器端读到数据", string(buf[:n]))
}
