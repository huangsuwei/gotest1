package main

import (
	"fmt"
	"net"
	"time"
)

func main()  {
	//组织一个udp地址结构，指定服务器的ip+port
	srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("net.resolve err", err)
		return
	}
	fmt.Println("服务器地质结构创建完成")


	//创建用于通信的socket
	udpConn, err1 := net.ListenUDP("udp", srvAddr)
	if err1 != nil {
		fmt.Println("net.listen err", err1)
		return
	}
	defer udpConn.Close()
	fmt.Println("通信socket创建完成")

	//读取客户端发送的数据
	buf := make([]byte, 4096)
	//返回三个数，字节数，客户端地址，错误信息
	n, cliAddr, err2 := udpConn.ReadFromUDP(buf)
	if err2 != nil {
		fmt.Println("udpconn.read err", err2)
		return
	}

	fmt.Printf("服务器读到客户端：%v 的数据，%s\n", cliAddr, string(buf[:n]))

	//回写数据给客户端
	daytime := time.Now().String()
	_, err3 := udpConn.WriteToUDP([]byte(daytime), cliAddr)
	if err3 != nil {
		fmt.Println("udp write err", err3)
		return
	}
}
