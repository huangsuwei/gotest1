package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main()  {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("格式为：go run xxx.go 文件名")
		return
	}

	//提取文件名
	path := list[1]
	//获取文件属性
	fileInfo, err := os.Stat(path)//文件访问的绝对路径
	if err != nil {
		fmt.Println("os.stat err", err)
		return
	}

	fileName := fileInfo.Name()

	//主动发起连接请求
	conn, err1 := net.Dial("tcp", "127.0.0.1:8008")
	if err1 != nil {
		fmt.Println("net.dial err", err1)
		return
	}
	defer conn.Close()

	//发送文件名给接收端
	conn.Write([]byte(fileName))

	//读取服务器回发的ok
	buf := make([]byte, 16)
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("conn.read err", err2)
		return
	}

	if string(buf[:n]) == "ok" {
		//写文件内容给服务器
		sendFile(conn, path)
	}
}

func sendFile(conn net.Conn, filePath string)  {
	//只读打开文件
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.open err", err)
		return
	}
	defer f.Close()

	//从本地文件中读取数据，写给接收端，读多少，写多少
	buf := make([]byte, 4096)
	for {
		n, err1 := f.Read(buf)
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println("发送文件完毕")
			} else {
				fmt.Println("f.read err", err1)
			}
			return
		}
		//写到网络socket中
		_, err2 := conn.Write(buf[:n])
		if err2 != nil {
			fmt.Println("conn.write err", err2)
		}
	}

}
