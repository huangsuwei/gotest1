package main

import (
	"fmt"
	"net/http"
)

func main() {
	//获取服务器应答包内容
	resp, err := http.Get("https://movie.douban.com/top250?start=0&filter=")
	if err != nil {
		fmt.Println("http.get err", err)
		return
	}
	defer resp.Body.Close()

	//简单查看应答包
	fmt.Println("header", resp.Header)
	fmt.Println("status", resp.Status)
	fmt.Println("statusCode", resp.StatusCode)
	fmt.Println("proto", resp.Proto)

	buf := make([]byte, 4096)
	var result string
	for {
		n, err1 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("-------read finish")
			break
		}
		if err1 != nil {
			fmt.Println("body.read.err", err1)
			break
		}
		result += string(buf[:n])
	}

	fmt.Printf("|%v|\n", result)
}
