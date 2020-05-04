package main

import (
	"fmt"
	"os"
)

func createFile(filePath string)  {
	fp, err := os.Create(filePath)
	if err != nil {
		fmt.Println("cuowu:", err)
	}

	//延时操作，在程序退出时，关闭文件
	//文件打开要及时关闭，不造成资源浪费，最大文件数65535
	defer fp.Close()
}

func main()  {
	filePath := "a.txt"
	createFile(filePath)
}
