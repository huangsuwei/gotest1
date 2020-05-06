package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.listen err", err)
		return
	}
	defer listener.Close()

	go func() {
		conn, err1 := listener.Accept()
		if err1 != nil {
			fmt.Println("accept err", err)
			return
		}
		defer conn.Close()

		conn.Write([]byte("hello world"))
	}()

	for {

	}
}
