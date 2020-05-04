package main

import (
	"fmt"
	"net"
)

func main()  {
	//指定服务器ip和port，创建通信套接字
	connect, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.dial err", err)
		return
	}
	defer connect.Close()

	//主动写数据给服务器
	connect.Write([]byte("are you ok"))

	buf := make([]byte, 4096)
	//接受服务器回发的数据
	n, err1 := connect.Read(buf)
	if err1 != nil {
		fmt.Println("服务器回发数据读取失败", err1)
		return
	}
	//读多少，写多少
	connect.Write(buf[:n])

	fmt.Println("服务器回发数据：", string(buf[:n]))
}
