package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "abc a7c mfc cat 8ca azc cba aMc"
	//解析，编译正则表达式
	//ret, err := regexp.Compile("a.c")//``:表示使用原生字符串
	ret, err := regexp.Compile(`a[^0-9a-z]c`) //``:表示使用原生字符串
	if err != nil {
		fmt.Println("regexp.compile err", err)
		return
	}
	//提取需要信息
	alls := ret.FindAllString(str, -1)
	fmt.Println("alls :", alls)
}
