package main

import (
	"fmt"
	"github.com/foxsuagr-sanse/go-gobang_game/common/utils"
	"reflect"
	"unsafe"
)

type MyCustomClaims struct {
	Name string
	Age  int64
	Cd   map[string]string
}

func main() {
	data := utils.UserInput(MyCustomClaims{Name: "ss"})
	if data != nil {
		println(data.Empty,data.Exists,data.FieldNumber)
	}
	mains()
}

func mains() {
	s := make([]int, 9, 20)
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(Len, len(s)) // 9 9

	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(Cap, cap(s)) // 20 20

	// map
	//var slcn interface{}
	slc := make([]string,0)
	//slcn = slc
	FMT(slc)
	var oldmap interface{}
	oldmap = make(map[string]interface{})
	m := oldmap.(map[interface{}]interface{})
	m["nmsl"] = "sb"
	count := **(**int)(unsafe.Pointer(&m))
	FMT(oldmap)
	FMT(count)
}

func FMT(v interface{}) {
	fmt.Println(reflect.TypeOf(v))
}