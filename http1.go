package main

import (
	"fmt"
	"net"
	"os"
)

func main()  {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.listen err", err)
		return
	}
	defer listener.Close()

	conn, err1 := listener.Accept()
	errFunc(err1, "listener.accept")
	defer conn.Close()

	buf := make([]byte, 4096)
	n, err2 := conn.Read(buf)
	if n == 0 {
		return
	}
	errFunc(err2, "conn.read")

	fmt.Printf("|%s|\n", string(buf[:n]))
}

func errFunc(err error, info string)  {
	if err != nil {
		fmt.Println(info + " err", err)
		//将当前进程结束
		os.Exit(1)
	}
}
