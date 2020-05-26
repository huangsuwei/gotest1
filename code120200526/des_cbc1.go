package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

//des 的CBC加密
//编写填充函数，如果最后一个分组字节不够，填充
//。。。。。。字节数刚好合适，添加一个新的分组
//填充的字节的值 == 缺少的字节的数
func PaddingLastGroup(plainText []byte, blockSize int) []byte {
	//1.求出最后一个组中剩余的字节数 28 % 8 = 3...4  32 % 8 = 4...0
	padNum := blockSize - len(plainText)%blockSize
	//2.创建新的切片，长度 == padNum,每个字节值byte(padNum)
	char := []byte{byte(padNum)} //长度1，
	//切片创建，并初始化
	newPlain := bytes.Repeat(char, padNum)
	//3.newPlain数组追加到原始铭文的后边

	return append(plainText, newPlain...)
}

//把填充的数据去掉
func unPaddingLastGroup(plainText []byte) []byte {
	//1.拿去切片中的最后一个字节
	length := len(plainText)
	lastChar := plainText[length-1]
	number := int(lastChar) //尾部填充的个数

	return plainText[:length-number]
}

//加密
func desEncrypt(plainText, key []byte) []byte {
	//1.建一个底层的des接口
	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	//2.明文填充
	newText := PaddingLastGroup(plainText, block.BlockSize())
	//3.创建一个使用cbc分组接口
	iv := []byte("12345678")
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//4.加密
	cipherText := make([]byte, len(newText))
	blockMode.CryptBlocks(cipherText, newText)

	return cipherText
}

//解密
func desDecrypt(cipherText, key []byte) []byte {
	//1.建一个底层的des接口
	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	//2.创建一个使用des解密的接口
	iv := []byte("12345678")
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//3.解密
	blockMode.CryptBlocks(cipherText, cipherText)
	//4.cipherText现在存储的是明文，需要删除加密时填充的尾部数据
	return unPaddingLastGroup(cipherText)
}

func main() {
	fmt.Println("des加密")
	key := []byte("1234abcd")
	src := []byte("如果w.W字段实现了io.Closer接口，本方法会调用其Close方法并返回该方法的返回值；否则不做操作返回nil。")
	cipherText := desEncrypt(src, key)
	fmt.Println("加密后的密文:", cipherText)
	plainText := desDecrypt(cipherText, key)
	fmt.Printf("解密后的明文:%s\n", string(plainText))
}
