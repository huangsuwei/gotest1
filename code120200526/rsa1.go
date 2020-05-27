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
	//4.pem编码
	file, err1 := os.Create("private.pem")
	if err1 != nil {
		fmt.Println(err1)
	}
	pem.Encode(file, &block)
	file.Close()

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
	//4.pem编码
	publicFile, err3 := os.Create("public.pem")
	if err3 != nil {
		fmt.Println(err3)
	}
	pem.Encode(publicFile, &publicBlock)
	publicFile.Close()
}

func main() {
	GenerateRsaKey(1024)
}
