package main

import (
	"fmt"
	"net"
	"os"
)

func main()  {
	//创建用于监听的socket
	listener, err1 := net.Listen("tcp", "127.0.0.1:8008")
	if err1 != nil {
		fmt.Println("net.listen err", err1)
		return
	}
	defer listener.Close()

	//阻塞监听
	conn, err2 := listener.Accept()
	if err2 != nil {
		fmt.Println("accept err", err2)
		return
	}
	defer conn.Close()

	//获取文件名，保存
	buf := make([]byte, 4096)
	n, err3 := conn.Read(buf)
	if err3 != nil {
		fmt.Println("conn.read err", err3)
		return
	}
	fileName := string(buf[:n])

	//回写ok给发送端
	conn.Write([]byte("ok"))

	//获取文件内容
	receFile(conn, fileName)
}

func receFile(conn net.Conn, fileName string)  {
	//按照文件名创建新文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.create err", err)
		return
	}
	defer f.Close()

	//从网络中读取数据写到本地
	buf := make([]byte, 4096)
	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			fmt.Println("文件接收完成")
			return
		}
		//读多少，写多少
		f.Write(buf[:n])
	}
}