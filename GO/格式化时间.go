package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println(time.Now().Format("2006-01-02"))
	fmt.Println(time.Now().Format("2006-01-02-03-04"))
	fmt.Println(time.Now().Format("2006-01-02-15-04"))
	fmt.Println(time.Now().Format("2006-01-01-15-04"))
	fmt.Println(time.Now().String())
}
