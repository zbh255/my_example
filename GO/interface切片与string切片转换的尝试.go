package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func toStringSlice(v interface{}) []string {
	sv := v.([]interface{})
	svHeader := (*reflect.SliceHeader)(unsafe.Pointer(&sv))
	strSlice := make([]string, 0)
	ssHeader := (*reflect.SliceHeader)(unsafe.Pointer(&strSlice))
	ssHeader.Cap = svHeader.Cap
	ssHeader.Len = svHeader.Len
	const i = 7
	svData := (*[i]interface{})(unsafe.Pointer(svHeader.Data))
	ssData := *(*[i]string)(unsafe.Pointer(svData))
	ssHeader.Data = (uintptr)(unsafe.Pointer(&ssData))
	return *(*[]string)(unsafe.Pointer(ssHeader))
}

func main() {
	n := make([]interface{}, 7)
	for i := 0; i < len(n); i++ {
		n[i] = "ll"
	}
	strSlice := toStringSlice(n)
	fmt.Println(strSlice)
}
