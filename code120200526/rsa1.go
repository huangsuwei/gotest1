package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

//生成rsa秘钥对，并保存到文件中
func GenerateRsaKey(keySize int) {
	//1.使用rsa中的generateKey函数生成秘钥
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		fmt.Println(err)
	}
	//2.通过x509标准将得到的私钥序列化为ASN.1的DER编码字符串
	derText := x509.MarshalPKCS1PrivateKey(privateKey)
	//3.要组织一个pem.block
	block := pem.Block{
		Type:  "rsa private key", //这个地方写个字符串就行
		Bytes: derText,
	}
	WriteKeyToFile("private.pem1", block)

	/**--------------------公钥----------------------**/
	//1.从私钥中，取出公钥
	publicKey := privateKey.PublicKey
	//2.使用x509标准序列化
	derstream, err2 := x509.MarshalPKIXPublicKey(&publicKey)
	if err2 != nil {
		fmt.Println(err2)
	}
	//3.将得到的数据放到block中
	publicBlock := pem.Block{
		Type:  "rsa public key",
		Bytes: derstream,
	}
	WriteKeyToFile("public.pem1", publicBlock)
}

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

func main() {
	//公钥钥长度和加密的数据时有关联的，如果数据太长了，下面这个也要相应的keySize也要相应变长
	GenerateRsaKey(4096)
	src := []byte("我是小崔，如果我死了，我肯定不是自杀...我是小崔，如果我死了，我肯定不是自杀...我是小崔，如果我死了，我肯定不是自杀...我是小崔，如果我死了，我肯定不是自杀...")
	cipherText := RsaEncrypt(src, "public.pem1")
	fmt.Println("加密后的密文：", cipherText)
	plainText := RsaDecrypt(cipherText, "private.pem1")
	fmt.Printf("解密后的明文:%s\n", string(plainText))
}

//rsa加密，公钥加密
func RsaEncrypt(plainText []byte, fileName string) []byte {
	//打开文件，读取文件内容
	buf := ReadFile(fileName)

	//2.pem解码
	block, _ := pem.Decode(buf)
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	pubKey, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		fmt.Println("断言失败")
	}

	//3.使用公钥加密
	cipherText, err3 := rsa.EncryptPKCS1v15(rand.Reader, pubKey, plainText)
	if err3 != nil {
		fmt.Println(err3)
	}

	return cipherText
}

//rsa解密，私钥解密
func RsaDecrypt(cipherText []byte, fileName string) []byte {
	//打开文件，读取文件内容
	buf := ReadFile(fileName)

	//2.pem解码
	block, _ := pem.Decode(buf)
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	//3.使用私钥解密
	plainText, err3 := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err3 != nil {
		fmt.Println(err3)
	}

	return plainText
}
