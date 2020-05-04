package main

import (
	"fmt"
	"os"
)

func main()  {
	list := os.Args

	if len(list) != 2 {
		fmt.Println("格式为：go run xxx.go 文件名")
		return
	}

	//提取文件名
	path := list[1]
	//获取文件属性
	fileInfo, err := os.Stat(path)//文件访问的绝对路径
	if err != nil {
		fmt.Println("os.stat err", err)
		return
	}

	fmt.Println("文件名：", fileInfo.Name())
	fmt.Println("文件大小：", fileInfo.Size())
}
