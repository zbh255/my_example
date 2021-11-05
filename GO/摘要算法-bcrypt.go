package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func main() {
	for cost := 10; cost <= 20; cost++ {
		start := time.Now()
		pwd2, _ := bcrypt.GenerateFromPassword([]byte("pa55w0rd"), cost)
		println(hex.EncodeToString(pwd2))
		pwd := sha512.New()
		pwd.Write([]byte("pa55w0rd"))
		passwd := hex.EncodeToString(pwd.Sum(nil))
		println(passwd)
		fmt.Printf("cost: %d, duration: %v\n", cost, time.Since(start))
	}
}
