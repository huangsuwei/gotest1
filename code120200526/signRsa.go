package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	src := []byte("我是小崔，如果我死了，我肯定不是自杀...我是小崔，如果我死了，我肯定不是自杀...我是小崔，如果我死了，我肯定不是自杀...我是小崔，如果我死了，我肯定不是自杀...")
	signText := SignatureRsa(src, "private.pem")
	bl := VerifyRsa(src, signText, "public.pem")
	fmt.Println("验证结果：", bl)
}

//rsa签名-私钥
func SignatureRsa(plainText []byte, fileName string) []byte {
	//1.打开私钥文件
	//2.读文件
	buf := ReadFile(fileName)

	//3.pem解码，得到pem.block的结构体变量
	block, _ := pem.Decode(buf)
	//4.x509将数据解析成私钥结构体 -》 得到了私钥
	privateKey, err2 := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err2 != nil {
		fmt.Println("x509.ParsePKCS1PrivateKey err", err2)
	}
	//5.创建一个hash对象，sha1/md5 -> sha512
	/*myHash := sha512.New()
	//6.给hash添加数据
	myHash.Write(plainText)
	//7.计算hash值
	hashText := myHash.Sum(nil)*/
	//8.使用rsa中的函数对散列值签名
	signText, err3 := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA512, CalculateHashSum(plainText))
	if err3 != nil {
		fmt.Println("rsa.SignPKCS1v15", err3)
	}

	return signText
}

func CalculateHashSum(plainText []byte) []byte {
	//创建hash接口
	myHash := sha512.New()
	//添加数据
	myHash.Write(plainText)
	//计算散列值
	return myHash.Sum(nil)
}

//rsa签名校验
func VerifyRsa(plainText, signText []byte, pubFileName string) bool {
	//1.打开公钥文件，读取
	/*buf := code2020ReadFile(pubFileName)*/
	buf := ReadFile(pubFileName)

	//2.pem解码，得到结构体
	block, _ := pem.Decode(buf)
	//3.使用x509对结构体中的数据进行解析，得到的一个接口
	pubInterface, err2 := x509.ParsePKIXPublicKey(block.Bytes)
	if err2 != nil {
		fmt.Println("x509.ParsePKIXPublicKey err", err2)
	}
	//4.类型断言，得到公钥结构体
	pubKey := pubInterface.(*rsa.PublicKey)
	//5.对原始消息进行hash运算，使用hash运算一致，得到散列值
	hashText := CalculateHashSum(plainText)

	//6.签名认证，rsa中的函数
	return rsa.VerifyPKCS1v15(pubKey, crypto.SHA512, hashText[:], signText) == nil
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
