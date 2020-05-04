package main

import (
	"fmt"
	"os"
)

func createFile2(filePath string)  {
	//文件存在会覆盖，不存在会创建
	fp, err := os.Create(filePath)
	if err != nil {
		fmt.Println("cuowu:", err)
		return
	}

	//写入操作
	n, err := fp.WriteString("hello, world!,传智播客")
	if err != nil {
		fmt.Println("cuowu ", err)
	}
	fmt.Println("写入字符", n)

	//延时操作，在程序退出时，关闭文件
	//文件打开要及时关闭，不造成资源浪费，最大文件数65535
	defer fp.Close()
}

func main()  {
	filePath := "a1.txt"
	createFile2(filePath)
}