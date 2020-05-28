package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	myHash()
}

func myHash() {
	//使用hash接口对象
	myHash := sha256.New()
	//添加数据
	src := []byte("我是小崔，如果我死了，我肯定不是自杀...我是小崔，如果我死了，我肯定不是自杀...我是小崔，如果我死了，我肯定不是自杀...我是小崔，如果我死了，我肯定不是自杀...")
	myHash.Write(src)
	myHash.Write(src)
	myHash.Write(src)
	//3.计算结果
	res := myHash.Sum(nil)
	//4.转成十六进制
	myStr := hex.EncodeToString(res)
	fmt.Println("原始返回数据：", res)
	fmt.Printf("十六进制：%s\n", myStr)

	//下面这种方式适用于少量数据
	size := sha256.Sum256(src)
	fmt.Println("返回数据:", size)
}
