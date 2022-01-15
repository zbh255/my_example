package main

import "fmt"

func Swap(x, y *int) {
	*y = *x ^ *y
	*x = *x ^ *y
	*y = *x ^ *y
}

func main() {
	x, y := 7, 9
	fmt.Println(x, y)
	Swap(&x, &y)
	fmt.Println(x, y)
}
