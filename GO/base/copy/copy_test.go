package copy

import (
	"reflect"
	"testing"
	"unsafe"
)

// 测试一些类型的拷贝行为

// go 1.17
// 测试证明go string并非深拷贝两个string header的data地址一致
// 测试函数调用时关闭内联优化
func TestStringTypeCopy(t *testing.T) {
	str1 := "isCopy?"
	str2 := str1

	data1 := (*reflect.StringHeader)(unsafe.Pointer(&str1)).Data
	data2 := (*reflect.StringHeader)(unsafe.Pointer(&str2)).Data
	if data1 == data2 {
		t.Log("string cop is not deep copy")
	}
	str3 := funcCall(str1)
	data3 := (*reflect.StringHeader)(unsafe.Pointer(&str3)).Data
	if data1 == data3 {
		t.Log("string cop is not deep copy")
	}
}

func funcCall(s string) string {
	return s
}

// go 1.17
// 测试证明go slice并非深拷贝两个slice header的data地址一致
func TestSliceTypeCopy(t *testing.T) {
	slice1 := []string{"hello", "world"}
	slice2 := slice1

	data1 := (*reflect.SliceHeader)(unsafe.Pointer(&slice1)).Data
	data2 := (*reflect.SliceHeader)(unsafe.Pointer(&slice2)).Data
	if data1 == data2 {
		t.Log("string cop is not deep copy")
	}
}
