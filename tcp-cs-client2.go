package main

import (
	"fmt"
	"net"
	"os"
)

func main()  {
	connect, err1 := net.Dial("tcp", "127.0.0.1:8001")
	if err1 != nil {
		fmt.Println("net.dial err", err1)
		return
	}
	defer connect.Close()

	//获取用户键盘输入stdin，将输入的数据发送给服务器
	go func() {
		str := make([]byte, 4096)
		for {
			n, err2 := os.Stdin.Read(str)
			if err2 != nil {
				fmt.Println("os.stdin.read err", err2)
				continue
			}
			//写给服务器，读多少，写多少
			connect.Write(str[:n])
		}
	}()

	//回显服务器返回的大写数据stdout
	buf := make([]byte, 4096)
	for {
		n2, err3 := connect.Read(buf)
		if n2 == 0 {
			fmt.Println("检测到服务器端关闭，客户端也关闭")
			return
		}
		if err3 != nil {
			fmt.Println("connect.read err", err3)
			return
		}
		fmt.Println("服务器读到客户端回发数据：", string(buf[:n2]))
	}
}
