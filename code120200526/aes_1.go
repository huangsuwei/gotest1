package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

//加密
func aesEncrypt(plainText, key []byte) []byte {
	//1.建一个底层的des接口
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	//2.不需要明文填充
	//newText := PaddingLastGroup(plainText, block.BlockSize())
	//3.创建一个使用ctr分组接口
	iv := []byte("12345678ABCDEFGH")
	stream := cipher.NewCTR(block, iv)
	//4.加密
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return cipherText
}

//解密
func aesDecrypt(cipherText, key []byte) []byte {
	//1.建一个底层的des接口
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	//2.创建一个使用des解密的接口
	iv := []byte("12345678ABCDEFGH")
	stream := cipher.NewCTR(block, iv)
	//3.解密
	stream.XORKeyStream(cipherText, cipherText)
	//blockMode.CryptBlocks(cipherText, cipherText)
	//4.cipherText现在存储的是明文，需要删除加密时填充的尾部数据
	return cipherText
}

func main() {
	fmt.Println("aes加密")
	key := []byte("abcdefgh12345678")
	src := []byte("如果w.W字段实现了io.Closer接口，本方法会调用其Close方法并返回该方法的返回值；否则不做操作返回nil。")
	cipherText := aesEncrypt(src, key)
	fmt.Println("加密后的密文:", cipherText)
	plainText := aesDecrypt(cipherText, key)
	fmt.Printf("解密后的明文:%s\n", string(plainText))
}
