package base

import (
	"fmt"
	"testing"
	"unsafe"
)

type S1 struct {
	b1 byte
	b2 byte
	b3 byte
	i1 int32
}

type S2 struct {
	s1 []byte
	b1 byte
	b3 byte
	i1 int32
	i2 int64
}

// 探讨Go Struct的内存对齐规则
// 部分代码思想来自《Go语言学习笔记》
func TestStructAlignment(t *testing.T) {
	var s1 = &S1{}
	var s2 = &S2{}
	s1Fmt := `
S1.b1 : %p , size: %d, offset: %d
S1.b2 : %p , size: %d, offset: %d
S1.b3 : %p , size: %d, offset: %d
S1.i1 : %p , size: %d, offset: %d
			`
	s2Fmt := `
S2.s1 : %p , size: %d, offset: %d
S2.b1 : %p , size: %d, offset: %d
S2.b3 : %p , size: %d, offset: %d
S2.i1 : %p , size: %d, offset: %d
S2.i2 : %p , size: %d, offset: %d
`
	fmt.Printf("ptr=%p,align=%d,size=%d\n", s1, unsafe.Alignof(s1), unsafe.Sizeof(*s1))
	fmt.Printf(s1Fmt, &s1.b1, unsafe.Sizeof(s1.b1), unsafe.Offsetof(s1.b1),
		&s1.b2, unsafe.Sizeof(s1.b2), unsafe.Offsetof(s1.b2),
		&s1.b3, unsafe.Sizeof(s1.b3), unsafe.Offsetof(s1.b3),
		&s1.i1, unsafe.Sizeof(s1.i1), unsafe.Offsetof(s1.i1))
	fmt.Printf("ptr=%p,align=%d,size=%d\n", s2, unsafe.Alignof(s2), unsafe.Sizeof(*s2))
	fmt.Printf(s2Fmt, &s2.s1, unsafe.Sizeof(s2.s1), unsafe.Offsetof(s2.s1),
		&s2.b1, unsafe.Sizeof(s2.b1), unsafe.Offsetof(s2.b1),
		&s2.b3, unsafe.Sizeof(s2.b3), unsafe.Offsetof(s2.b3),
		&s2.i1, unsafe.Sizeof(s2.i1), unsafe.Offsetof(s2.i1),
		&s2.i2, unsafe.Sizeof(s2.i2), unsafe.Offsetof(s2.i2))
}
