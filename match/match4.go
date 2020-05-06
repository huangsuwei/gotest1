package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var start, end int
	fmt.Print("请输入爬取的起始页（>= 1）")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的结束页（>= start）")
	fmt.Scan(&end)

	toWork2(start, end)
}

func toWork2(start, end int) {
	fmt.Printf("正在爬取%d到%d页的数据", start, end)

	page := make(chan int)

	for i := start; i <= end; i++ {
		go SpiderPage3(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d页数据存储完成！", <-page)
	}
}

func SpiderPage3(idx int, page chan int) {
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(idx) + ".html"

	//封装函数获取段子的url
	result, err := httpGet3(url)
	if err != nil {
		fmt.Println("httpget3.err", err)
		return
	}

	//解析，编译正则
	ret1, _ := regexp.Compile(`<h1 class="dp-b"><a href="(?s:(.*?))`)
	//提取需要信息--每个段子的url
	alls := ret1.FindAllStringSubmatch(result, -1)

	//创建用户存储的切片
	fileTitle := make([]string, 0)
	fileContent := make([]string, 0)

	for _, jokeUrl := range alls {
		title, content, err1 := SpiderJokePage(jokeUrl[1])
		if err1 != nil {
			fmt.Println("spiderJokePage err", err1)
			continue
		}

		fileTitle = append(fileTitle, title)
		fileContent = append(fileContent, content)
	}

	saveJokeToFile(idx, fileTitle, fileContent)

	page <- idx
}

func saveJokeToFile(idx int, fileTitle, fileContent []string) {
	path := "D:/gowork/gotest1/files/jokes/" + "第" + strconv.Itoa(idx) + "页.txt"
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("os.create err", err)
		return
	}
	defer f.Close()

	n := len(fileTitle)
	for i := 0; i < n; i++ {
		f.WriteString(fileTitle[i] + "\n" + fileContent[i] + "\n")
		f.WriteString("---------------------------------------\n")
	}
}

func SpiderJokePage(url string) (title, content string, err error) {
	result, err1 := httpGet3(url)
	if err1 != nil {
		err = err1
		return
	}
	//编译解析正则表达式
	ret1, _ := regexp.Compile(`<h1>(?s:(.*?))</h1>`) //有两处，取第一个
	alls := ret1.FindAllStringSubmatch(result, -1)
	for _, tempTitle := range alls {
		title = tempTitle[1]
		title = strings.Replace(title, " ", "", -1)
		title = strings.Replace(title, "\n", "", -1)
		title = strings.Replace(title, "\t", "", -1)
		break
	}

	//编译解析正则表达式
	ret2, _ := regexp.Compile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`) //有两处，取第一个
	alls2 := ret2.FindAllStringSubmatch(result, -1)
	for _, tempContent := range alls2 {
		content = tempContent[1]
		content = strings.Replace(content, " ", "", -1)
		content = strings.Replace(content, "\n", "", -1)
		content = strings.Replace(content, "&nbsp;", "", -1)
		break
	}

	return
}

func httpGet3(url string) (result string, err error) {
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
