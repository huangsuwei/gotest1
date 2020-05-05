package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	//指定爬取起始，终止页
	var start, end int
	fmt.Print("请输入爬取的起始页（》=1）：")
	fmt.Scan(&start)

	fmt.Print("请输入爬取的终止页（》=start）：")
	fmt.Scan(&end)

	working2(start, end)
}

func working2(start, end int) {
	fmt.Printf("正在爬取第%d到%d页的数据\n", start, end)

	page := make(chan int)

	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}
}

func HttpGet2(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1 //将封装函数内部的错误，传出调用者
		return
	}
	defer resp.Body.Close()

	//循环读取网页数据，穿出给调用者
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			//fmt.Println("网页读取完成")
			break
		}

		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}

	return result, err
}

//爬单个页面的数据
func SpiderPage(i int, page chan int) {
	//循环爬取每一页的数据
	url := "https://tieba.baidu.com/f?ie=utf-8&kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&pn=" + strconv.Itoa((i-1)*50)
	result, err := HttpGet2(url)
	if err != nil {
		fmt.Println("httpGet err", err)
		return
	}
	//fmt.Println("result=", result)

	//将读到的整网页数据保存成一个文件
	f, err1 := os.Create("第" + strconv.Itoa(i) + "页" + ".html")
	if err1 != nil {
		fmt.Println("os.create err", err1)
		return
	}
	f.WriteString(result)

	f.Close() //保存好一个文件，关闭一个文件

	page <- i
}
