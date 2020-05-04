package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(srcFile string, destFile string)  {
	if srcFile == destFile {
		fmt.Println("原文件与目标文件不得重名")
	}
	//打开文件
	sfp, err1 := os.Open(srcFile)
	if err1 != nil {
		fmt.Println("cuowu", err1)
	}

	dfp, err2 := os.Create(destFile)
	if err2 != nil {
		fmt.Println("cuowu", err2)
	}
	//原文件内容拷贝到新文件
	buf := make([]byte, 1024 * 8)//8k
	for {
		n, err := sfp.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("cuowu", err)
		}
		dn, err3 := dfp.Write(buf[:n])
		fmt.Println(dn, err3)
	}

	//关闭文件
	defer sfp.Close()
	defer dfp.Close()
}

func main()  {
	copyFile("a1.txt", "a.txt")
}
