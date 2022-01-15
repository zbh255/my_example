package main

import "fmt"

func main() {
	x1 := uint8(0x63)
	x2 := uint8(0x95)
	fmt.Printf("左移4位:%b\n", x1<<4)
	fmt.Printf("左移4位:%b\n", x2<<4)
	fmt.Printf("%d\n", 99<<4)
	fmt.Printf("%b\n", x1>>4)
	fmt.Printf("%b", x2>>4)
}
