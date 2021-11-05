package main

import (
	"fmt"
	this "github.com/mengzushan/bups/common/error"
	"github.com/mengzushan/bups/common/logger"
	"reflect"
)

func main() {
	_, err := logger.Std("./sudo/sudi/app.log")
	defer Error()
	if err != this.Nil {
		panic(err)
	}
}

func Error() {
	err := recover()
	fmt.Println(reflect.TypeOf(err), reflect.TypeOf(this.LogError))
	switch err.(type) {
	case this.Error:
		panic(err)
	}
}
