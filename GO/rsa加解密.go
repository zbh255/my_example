package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	NewKey()
	Use()
}

func NewKey() {
	// 生成私钥
	privatekey, err := rsa.GenerateKey(rand.Reader,1024)
	if err != nil {
		panic(err)
	}

	// 使用x509标准转换为.pem格式的文件
	derText := x509.MarshalPKCS1PrivateKey(privatekey)
	
	// 组织一个pem.block
	block := pem.Block{
		Type: "rsa private key",
		Bytes: derText,
	}
	pathHead, _ := os.Getwd()
	// pem编码
	file,err2 := os.Create(pathHead + "/cache/rsa/private.pem")
	if err2 != nil {
		panic(err2)
	}
	_ = pem.Encode(file, &block)
	defer file.Close()

	// 获得私钥对应的公钥
	publickey := privatekey.PublicKey

	// x509序列化
	derpText, err3 := x509.MarshalPKIXPublicKey(&publickey)
	if err3 != nil {
		panic(err3)
	}

	// 组织一个pem.block
	block = pem.Block{
		Type: "rsa public key",
		Bytes: derpText,
	}

	// pem编码
	file2,err4 := os.Create(pathHead + "/cache/rsa/public.pem")
	if err4 != nil {
		panic(err4)
	}
	_ = pem.Encode(file2, &block)
	defer file2.Close()
}

func Use() {
	// 使用公钥加密,私钥解密
	// 打开公钥文件
	pathHead, _ := os.Getwd()
	file,err := os.Open(pathHead + "/cache/rsa/public.pem")
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println(fileInfo.Size())
	buf := make([]byte,fileInfo.Size())
	_, _ = file.Read(buf)
	defer file.Close()

	// pem解码
	block, _ := pem.Decode(buf)

	// 使用x509标准转换成可以使用的公钥
	pk, _ := x509.ParsePKIXPublicKey(block.Bytes)

	// 强制转换
	publicKey := pk.(*rsa.PublicKey)

	// 使用公钥加密数据
	cipherText,err := rsa.EncryptPKCS1v15(rand.Reader,publicKey, []byte("Hello, world!"))
	if err != nil {
		panic(err)
	}
	files, _ := os.Create(pathHead + "/cache/rsa/key.txt")
	_, _ = files.Write(cipherText)
	defer files.Close()
	fmt.Println("加密后: " + hex.EncodeToString(cipherText))
	fmt.Println("加密后: " + base64.StdEncoding.EncodeToString(cipherText))

	src, _ := base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString(cipherText))
	// 私钥解密
	// 打开私钥文件
	prikeyfile,err := os.Open(pathHead + "/cache/rsa/private.pem")
	if err != nil {
		panic(err)
	}
	defer prikeyfile.Close()
	fileInfo,err = prikeyfile.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println(fileInfo.Size())
	buf = make([]byte,fileInfo.Size())
	_, _ = prikeyfile.Read(buf)

	// pem解密
	block, _ = pem.Decode(buf)
	privatekey,err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	// 使用私钥解密密文
	plainText,err := rsa.DecryptPKCS1v15(nil,privatekey,cipherText)
	if err != nil {
		panic(err)
	}
	plainText2,err := rsa.DecryptPKCS1v15(nil,privatekey,src)
	if err != nil {
		panic(err)
	}
	fmt.Println("密文: " + string(plainText))
	fmt.Println("密文: " + string(plainText2))
}