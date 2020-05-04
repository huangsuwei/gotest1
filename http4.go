package main

import (
	"fmt"
	"net/http"
	"os"
)

func main()  {
	//注册回调函数，该函数在客户端访问服务器时，会被自动调用
	http.HandleFunc("/", myHandle)

	//绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8000", nil)
}


func myHandle(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("客户端请求：", r.URL)
	openSendFile(r.URL.String(), w)
}

func openSendFile(fName string, w http.ResponseWriter)  {
	pathFileName := "/" + fName
	f, err := os.Open(pathFileName)
	if err != nil {
		fmt.Println("open err", err)
		w.Write([]byte("no such file or directory !"))
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		//从本地将文本内容读取
		n, _ := f.Read(buf)
		if n == 0 {
			return
		}
		w.Write(buf[:n])
	}
}