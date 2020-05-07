package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {
	url := "https://www.douyu.com/g_yz"

	//爬取整个页面的全部信息存储在result中
	result, err := httpGet4(url)
	if err != nil {
		fmt.Println("http get err", err)
		return
	}

	//解析编译正则
	ret1, _ := regexp.Compile(`<img loading="lazy" src="(?s:(.*?))"`)
	//提取需要信息--每个段子的url
	alls := ret1.FindAllStringSubmatch(result, 50)

	page := make(chan int)
	n := len(alls)
	for idx, imgUrl := range alls {
		go saveImg(idx, imgUrl[1], page)
	}

	for i := 0; i < n; i++ {
		//runtime.GC()//防止主go程退出
		fmt.Printf("第%d张图片下载完成\n", <-page+1)
	}
}

func saveImg(idx int, imgUrl string, page chan int) {
	path := "D:/gowork/gotest1/files/douyu/" + "第" + strconv.Itoa(idx) + "number" + strconv.Itoa(rand.Int()) + "页.jpg"
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("os.create err", err)
		return
	}
	defer f.Close()

	resp, err1 := http.Get(imgUrl)
	if err1 != nil {
		fmt.Println("http.get err", err1)
		return
	}

	defer resp.Body.Close()

	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Printf("第%d张图片读取完成！\n", idx+1)
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		f.Write(buf[:n])
	}

	page <- idx
}

func httpGet4(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取完成!")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}

		result += string(buf[:n])
	}

	return
}
