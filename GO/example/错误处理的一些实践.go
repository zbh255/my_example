package main

import (
	"fmt"
)

type ioError struct{ text string }

func (ioe *ioError) Error() string { return ioe.text }

type ERROR int

func (e ERROR) Error() string {
	return ""
}

const Nil ERROR = 1

var IOERROR = &ioError{text: "Invalid"}

func main() {
	err2 := err()
	if err2 == IOERROR {
		fmt.Println(err2)
	}
	println("world", Nil)
}

func err() error {
	return IOERROR
}
