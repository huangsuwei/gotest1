package main

import (
	"fmt"
	"strings"
)

func main()  {
	var str string = "hello world"
	//字符串查找
	value := strings.Contains(str, "world")
	fmt.Println(value)
	//字符串拼接
	strs := []string{"nihao", "nizaiganshenme", "gaoci"}
	buf := strings.Join(strs, "/")
	fmt.Println(buf)
	//字符串位置查找
	idx := strings.Index(str, "word")
	if idx == -1 {
		fmt.Println("未找到数据")
	} else {
		fmt.Println(idx)
	}
	//字符串重复
	str3 := strings.Repeat(str, 3)
	fmt.Println(str3)
	//字符串替换
	str4 := strings.Replace(str, "l", "", 1)
	fmt.Println(str4)
	//字符串分割,分割后是一个切片
	var str5 string = "hello@world@go"
	str6 := strings.Split(str5, "@")
	fmt.Println(str6)
	//去掉指定内容
	var str7 string = " ++       are you ok?   "
	str8 := strings.Trim(str7, "+")
	fmt.Println(str8)
	//字符串去空格,出来的是切片
	var str9 string = "     are you ok"
	str10 := strings.Fields(str9)
	for i, data := range str10 {
		fmt.Println(i, ",", data)
	}
	//fmt.Println(str10)
}
