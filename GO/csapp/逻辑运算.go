package main

import "fmt"

func main()  {
	a := uint8(0x66)
	b := uint8(0x39)
	fmt.Println(a&(-a))
	fmt.Println(^a|^b)
}
