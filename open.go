package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFile(fileName string)  {
	fp, err := os.Open(fileName)
	if err != nil {
		fmt.Println("cuowu", err)
		return
	}

	//定义切片读取内容
	buf := make([]byte, 1024)
	n, err1 := fp.Read(buf)
	//io.eof表示文件结尾
	if err1 != nil && err1 != io.EOF {
		fmt.Println("cuowu", err1)
		return
	}

	fmt.Println(string(buf[:n]))
	defer fp.Close()
}

func readFileLine(filePath string)  {
	fp, err := os.Open(filePath)
	if err != nil {
		fmt.Println("cuowu", err)
		return
	}

	//新建文件缓冲区，将数据放进去
	r := bufio.NewReader(fp)
	for {
		buf, err1 := r.ReadBytes('\n')
		fmt.Printf("buf = ###%s###", string(buf))
		if err1 != nil {
			if err1 != io.EOF {
				break
			}
			fmt.Println("cuowu", err1)
		}

	}

	defer fp.Close()
}

func main()  {
	fileName := "a1.txt"
	readFile(fileName)
}
