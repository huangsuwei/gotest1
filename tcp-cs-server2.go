package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func main()  {
	listener, err1 := net.Listen("tcp", "127.0.0.1:8001")
	if err1 != nil {
		fmt.Println("net.listen err", err1)
		return
	}
	defer listener.Close()

	for {
		fmt.Println("服务器等待客户端连接")
		connect, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println("accept err", err2)
			return
		}

		//具体完成客户端和服务器的
		go HandleConnect(connect)
	}
}

func HandleConnect(connect net.Conn)  {
	defer connect.Close()
	//获取客户端的地址
	addr := connect.RemoteAddr()
	fmt.Println(addr, "客户端连接成功！")

	buf := make([]byte, 4096)
	for {
		n, err := connect.Read(buf)
		if string(buf[:n]) == "exit\n" || string(buf[:n]) == "exit\r\n" {//会自带一个\n的字符!!!!!!!!!!!!!!!!!!
			fmt.Println("服务器接收到客户端关闭请求")
			return
		}

		if n == 0 {
			if err != nil && err == io.EOF {
				fmt.Println("客户端已与服务器断开连接")
			} else {
				fmt.Println("connect.read err", err)
			}
			return
		}
		if err != nil {
			fmt.Println("connect.read err", err)
			return
		}
		fmt.Println("服务器读到数据", string(buf[:n]))

		//完成小写转大写，回发给客户端
		connect.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}
