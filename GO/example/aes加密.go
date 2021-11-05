package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {
	// 明文
	d := []byte("hello,aes")
	key := []byte("1234567890/*-+12")
	fmt.Println("加密前: ", string(d))
	fmt.Println("加密前: ", hex.EncodeToString(d))
	b, err := hex.DecodeString(hex.EncodeToString(d))
	fmt.Println("加密前: ", string(b))
	x1, err := encryptAes(d, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("加密后: ", hex.EncodeToString(x1))
	// 解密
	x2, err := decryptAes(x1, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密后: ", string(x2))
}

// 加密函数
func encryptAes(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	src = padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src, nil
}

// 填充数据
func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}

// 解密函数
func decryptAes(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = unpadding(src, block.BlockSize())
	return src, nil
}

// 去除填充数据
func unpadding(src []byte, blockSize int) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	return src[:n-unPadNum]
}
