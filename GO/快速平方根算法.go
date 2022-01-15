package main

import (
	"fmt"
	"math"
	"unsafe"
)

/*
	卡马克算法
*/

func Qrsqrt(number float32) float32 {
	var i int
	var x2, y float32
	const threeHalfs float32 = 1.5

	x2 = number * 0.5
	y = number
	i = *(*int)(unsafe.Pointer(&y))
	i = 0x5f3759df - (i >> 1)
	y = *(*float32)(unsafe.Pointer(&i))
	y = y * (threeHalfs - (x2 * y * y))
	return y * number
}

func main() {
	fmt.Println(math.Sqrt(5))
	fmt.Print(Qrsqrt(5))
}
