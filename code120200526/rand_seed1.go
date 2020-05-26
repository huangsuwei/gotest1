package main

import (
	"fmt"
	"math/rand"
	"time"
)

//src -> 要加密的铭文, key -> 秘钥  大小为：8byte
/*func DesEncrypt_CBC(src, key []byte) []byte {
	//1.创建并返回一个使用DES算法的cipher.Block接口
	block, err := des.NewCipher(key)
	//2.判断是否创建成功
	if err != nil {
		panic(err)
	}
	//3.对最后一个明文分组进行数据填充
	src = PKCS5Padding(src, block.BlockSize())
	//4.创建一个密码分组为链接模式的，底层使用DES加密的BlockMode接口
	//参数iv的长度, 必须等于b的块尺寸
	tmp := []byte("helloAAA")
	blackMode := cipher.NewCBCEncrypter(block, tmp)
	//5.加密连续的数据块
	dst := make([]byte, len(src))
	blackMode.CryptBlocks(dst, src)
	fmt.Println("加密之后的数据：", dst)
}*/

func main() {
	rand.Seed(time.Now().UnixNano())

	randNum := rand.Int()

	fmt.Println(randNum)
}
