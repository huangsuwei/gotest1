package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var start, end int
	fmt.Print("请输入爬取的起始页（>= 1）")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的结束页（>= start）")
	fmt.Scan(&end)

	toWork(start, end)
}

func toWork(start, end int) {
	fmt.Printf("正在爬取%d到%d页的数据", start, end)

	page := make(chan int)

	for i := start; i <= end; i++ {
		go SpiderPage2(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d页数据存储完成！", <-page)
	}
}

func SpiderPage2(idx int, page chan int) {
	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((idx-1)*25) + "&filter="

	fmt.Println("网页为：", url)
	//封装httpget2获取数据
	result, err := httpGet2(url)
	fmt.Println("result=", result)
	if err != nil {
		fmt.Println("httpget.err", err)
		return
	}

	//解析，编译正则表达式---电影名称
	ret1, err := regexp.Compile(`<img width="100" alt="(?s:(.*?))"`)
	//提取需要信息
	fileName := ret1.FindAllStringSubmatch(result, -1)
	for _, name := range fileName {
		fmt.Println("name:", name[1])
	}

	//解析，编译正则表达式---电影名称
	ret2, err := regexp.Compile(`<span class="rating_nums">(?s:(.*?))</span>`)
	//提取需要信息
	score := ret2.FindAllStringSubmatch(result, -1)
	for _, name := range score {
		fmt.Println("score:", name[1])
	}

	//解析，编译正则表达式---电影名称
	ret3, err := regexp.Compile(`<span class="pl">((?s:(.*?))人评价)</span>`)
	//提取需要信息
	count := ret3.FindAllStringSubmatch(result, -1)
	for _, name := range count {
		fmt.Println("count:", name[1])
	}

	//将提取的有用信息封装到文件中
	save2File(idx, fileName, score, count)

	page <- idx
}

func httpGet2(url string) (result string, err error) {
	fmt.Println("读取网页：", url)
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	fmt.Println("爬取数据状态码：", resp.Status)

	buf := make([]byte, 4096)
	//循环爬取整页数据
	for {
		n, err2 := resp.Body.Read(buf)
		fmt.Println("读取数据：", n)
		if n == 0 {
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

func save2File(idx int, fileName, score, peopleNum [][]string) {
	path := "F:/gowork/gotest1/" + "第" + strconv.Itoa(idx) + "页.txt"

	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer f.Close()

	//先打印抬头， 电影名称  评分 评分人数
	n := len(fileName) //得到条目数，应该是25
	f.WriteString("电影名称" + "\t\t\t" + "评分" + "\t\t" + "评分人数" + "\n")
	for i := 0; i < n; i++ {
		f.WriteString(fileName[i][1] + "\t\t\t" + score[i][1] + "\t\t" + peopleNum[i][1] + "\n")
	}
}
