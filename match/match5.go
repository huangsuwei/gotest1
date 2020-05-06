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
	var start, end int
	/*fmt.Print("请输入爬取的起始页（>= 1）")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的结束页（>= start）")
	fmt.Scan(&end)*/
	start, end = 1, 1

	toWork3(start, end)
}

func toWork3(start, end int) {
	fmt.Printf("正在爬取%d到%d页的数据", start, end)

	page := make(chan int)

	for i := start; i <= end; i++ {
		go SpiderPage4(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d页数据存储完成！", <-page)
	}
}

func SpiderPage4(idx int, page chan int) {
	//url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(idx) + ".html"
	url := "https://www.douyu.com/g_yz"

	//封装函数获取段子的url
	result, err := httpGet4(url)
	if err != nil {
		fmt.Println("httpget4.err", err)
		return
	}

	//解析，编译正则
	//ret1, _ := regexp.Compile(`data-original="(?s:(.*?))`)
	ret1, _ := regexp.Compile(`<img loading="lazy" src="(?s:(.*?))"`)
	//提取需要信息--每个段子的url
	alls := ret1.FindAllStringSubmatch(result, 50)

	fmt.Println("alls", alls)

	count := 0
	for _, imgUrl := range alls {
		go saveImg(count, imgUrl[1], page)
		//fmt.Println("img:", imgUrl[1])
		count += 1
	}

	num := count
	for {
		switch <-page {
		case num:
			break
		}
	}
}

func saveImg(count int, imgUrl string, page chan int) {
	path := "D:/gowork/gotest1/files/douyu/" + "第" + strconv.Itoa(rand.Int()) + "页.jpg"
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
			fmt.Println("读取完成!")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		f.Write(buf[:n])
	}
	page <- count
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
