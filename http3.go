package main

import (
	"fmt"
	"net/http"
)

func main()  {
	//注册回调函数，该函数在客户端访问服务器时，会被自动调用
	http.HandleFunc("/itcast", myHandle)

	//绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func checkAllow(url string) bool {
	isAllow := false

	for _, allowUrl := range AllowUrls() {
		if url == allowUrl {
			isAllow = true
		}
	}

	return isAllow
}

func AllowUrls() map[int]string {
	var urls map[int]string

	urls[0] = "/itcast"

	return urls
}

func myHandle(w http.ResponseWriter, r *http.Request)  {
	//判断是否在允许访问的地址
	if !checkAllow(r.URL.String()) {
		fmt.Println("不允许访问，谢谢")
		return
	}

	//写给客户端的数据内容

	w.Write([]byte("this is a webServer"))
	//r从客户端读到的内容
	fmt.Println("Header:", r.Header)
	fmt.Println("Url:", r.URL)
	fmt.Println("Method:", r.Method)
	fmt.Println("Host:", r.Host)
	fmt.Println("RemoteAddr:", r.RemoteAddr)
	fmt.Println("Body:", r.Body)
}
