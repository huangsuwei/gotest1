package myfile

import (
	"encoding/pem"
	"fmt"
	"os"
)

func WriteKeyToFile(fileName string, b pem.Block) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.create err", err)
		return
	}
	//4.pem编码
	pem.Encode(file, &b)
	file.Close()
}

//读文件
func ReadFile(fileName string) []byte {
	//打开私钥文件
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("os.open err", err)
	}

	//读文件
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	file.Close()

	return buf
}
