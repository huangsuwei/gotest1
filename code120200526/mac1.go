package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

func main() {
	src := []byte("在消息秘钥中，需要发送者和接受者共同持有一个秘钥，秘钥不能被攻击者知道")
	key := []byte("hello world")
	hashValue := GenerateHMac(src, key)
	fmt.Println("验证结果：", VerifyHmac(src, key, hashValue))
}

//生成消息验证码函数
func GenerateHMac(plainText, key []byte) []byte {
	//1.创建hash接口，要指定使用的hash算法和秘钥
	myHash := hmac.New(sha1.New, key)
	//2.给hash对象添加数据
	myHash.Write(plainText)

	//计算散列值
	return myHash.Sum(nil)
}

//验证消息验证码
func VerifyHmac(plainText, key, hashText []byte) bool {
	hashValue := GenerateHMac(plainText, key)
	//1.创建hash接口，要指定使用的hash算法和秘钥
	/*myHash := hmac.New(sha1.New, key)
	//2.给hash对象添加数据
	myHash.Write(plainText)

	//比对散列值
	hMac := myHash.Sum(nil)*/

	return hmac.Equal(hashValue, hashText)
}
