package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)


func NewUint64() *Uint64 {
	return &Uint64{}
}

type Uint64 struct {
	buf uint64
}

func (n *Uint64) Load() []byte {
	return *(*[]byte)(unsafe.Pointer(uintptr(atomic.LoadUint64(&n.buf))))
}

func (n *Uint64) Store() {
	n.appendBuf()
}

func (n *Uint64) appendBuf() {
	tmp := "hello world"
	atomic.StoreUint64(&n.buf,(uint64)(uintptr(unsafe.Pointer(&tmp))))
}

/*
	TODO: 关于atomic.LoadUintptr & atomic.StoreUintptr未研究清楚的问题
	- 虽然将指针保存为uintptr，指向的区域可能会被gc回收，但是在escapes heap和通过GOGC=off关闭gc
	的情况下观察到以下情况
	-gcflags "-N -l" 关闭优化和内联，编译程序观察到b&b2的len&cap范围是一个很大数
	-gcflags "-l" 关闭内联，编译程序观察到b的len是正常的范围，cap是一个很大的数，b2的len&cap是一个很大的数
	不关闭内联或者优化的情况下程序一切正常
	- 猜测可能是编译器做了死码消除的优化使程序的行为看起来很正常，试图分析不同选项生成的汇编代码，但以我目前的
	水平看不出编译器做了什么优化。

	-gcflags "-N -l"下的Load函数
	TEXT main.(*Uint64).Load(SB) /Users/harder/Desktop/Git-Repo/github.com/abingzo/my_example/GO/base/stdlib/main.go
	main.go:18		0x108a180		4883ec30		SUBQ $0x30, SP
	main.go:18		0x108a184		48896c2428		MOVQ BP, 0x28(SP)
	main.go:18		0x108a189		488d6c2428		LEAQ 0x28(SP), BP
	main.go:18		0x108a18e		4889442438		MOVQ AX, 0x38(SP)
	main.go:18		0x108a193		48c744241000000000	MOVQ $0x0, 0x10(SP)
	main.go:18		0x108a19c		440f117c2418		MOVUPS X15, 0x18(SP)
	main.go:19		0x108a1a2		488b542438		MOVQ 0x38(SP), DX
	main.go:19		0x108a1a7		8402			TESTB AL, 0(DX)
	main.go:19		0x108a1a9		4889542408		MOVQ DX, 0x8(SP)
	main.go:19		0x108a1ae		488b12			MOVQ 0(DX), DX
	main.go:19		0x108a1b1		48891424		MOVQ DX, 0(SP)
	main.go:19		0x108a1b5		8402			TESTB AL, 0(DX)
	main.go:19		0x108a1b7		488b02			MOVQ 0(DX), AX
	main.go:19		0x108a1ba		488b5a08		MOVQ 0x8(DX), BX
	main.go:19		0x108a1be		488b4a10		MOVQ 0x10(DX), CX
	main.go:19		0x108a1c2		4889442410		MOVQ AX, 0x10(SP)
	main.go:19		0x108a1c7		48895c2418		MOVQ BX, 0x18(SP)
	main.go:19		0x108a1cc		48894c2420		MOVQ CX, 0x20(SP)
	main.go:19		0x108a1d1		488b6c2428		MOVQ 0x28(SP), BP
	main.go:19		0x108a1d6		4883c430		ADDQ $0x30, SP
	main.go:19		0x108a1da		c3			RET
	:-1			0x108a1db		cc			INT $0x3
	:-1			0x108a1dc		cc			INT $0x3
	:-1			0x108a1dd		cc			INT $0x3
	:-1			0x108a1de		cc			INT $0x3
	:-1			0x108a1df		cc			INT $0x3

	-gcflags "-l"下的Load函数
	TEXT main.(*Uint64).Load(SB) /Users/harder/Desktop/Git-Repo/github.com/abingzo/my_example/GO/base/stdlib/main.go
  	main.go:19		0x108a160		488b10			MOVQ 0(AX), DX
  	main.go:19		0x108a163		488b02			MOVQ 0(DX), AX
  	main.go:19		0x108a166		488b5a08		MOVQ 0x8(DX), BX
  	main.go:19		0x108a16a		488b4a10		MOVQ 0x10(DX), CX
  	main.go:19		0x108a16e		c3			RET
  	:-1			0x108a16f		cc			INT $0x3
  	:-1			0x108a170		cc			INT $0x3
  	:-1			0x108a171		cc			INT $0x3
  	:-1			0x108a172		cc			INT $0x3
  	:-1			0x108a173		cc			INT $0x3
  	:-1			0x108a174		cc			INT $0x3
  	:-1			0x108a175		cc			INT $0x3
  	:-1			0x108a176		cc			INT $0x3
  	:-1			0x108a177		cc			INT $0x3
  	:-1			0x108a178		cc			INT $0x3
  	:-1			0x108a179		cc			INT $0x3
  	:-1			0x108a17a		cc			INT $0x3
  	:-1			0x108a17b		cc			INT $0x3
  	:-1			0x108a17c		cc			INT $0x3
  	:-1			0x108a17d		cc			INT $0x3
  	:-1			0x108a17e		cc			INT $0x3
  	:-1			0x108a17f		cc			INT $0x3

	-gcflags "" 下的Load函数被内联

	-gcflags "-N -l" 下的appendBuf函数
	TEXT main.(*Uint64).appendBuf(SB) /Users/harder/Desktop/Git-Repo/github.com/abingzo/my_example/GO/base/stdlib/main.go
  	main.go:26		0x108a220		4883ec20		SUBQ $0x20, SP
  	main.go:26		0x108a224		48896c2418		MOVQ BP, 0x18(SP)
  	main.go:26		0x108a229		488d6c2418		LEAQ 0x18(SP), BP
  	main.go:26		0x108a22e		4889442428		MOVQ AX, 0x28(SP)
  	main.go:27		0x108a233		488d058b8a0100		LEAQ go.string.*+4133(SB), AX
  	main.go:27		0x108a23a		4889442408		MOVQ AX, 0x8(SP)
  	main.go:27		0x108a23f		48c74424100b000000	MOVQ $0xb, 0x10(SP)
  	main.go:28		0x108a248		488b442428		MOVQ 0x28(SP), AX
  	main.go:28		0x108a24d		8400			TESTB AL, 0(AX)
  	main.go:28		0x108a24f		48890424		MOVQ AX, 0(SP)
  	main.go:28		0x108a253		488d4c2408		LEAQ 0x8(SP), CX
  	main.go:28		0x108a258		488708			XCHGQ CX, 0(AX)
  	main.go:29		0x108a25b		488b6c2418		MOVQ 0x18(SP), BP
  	main.go:29		0x108a260		4883c420		ADDQ $0x20, SP
  	main.go:29		0x108a264		c3			RET

	-gcflags "-l" 下的appendBuf函数
	TEXT main.(*Uint64).appendBuf(SB) /Users/harder/Desktop/Git-Repo/github.com/abingzo/my_example/GO/base/stdlib/main.go
	  main.go:26		0x108a1c0		4883ec18		SUBQ $0x18, SP
	  main.go:26		0x108a1c4		48896c2410		MOVQ BP, 0x10(SP)
	  main.go:26		0x108a1c9		488d6c2410		LEAQ 0x10(SP), BP
	  main.go:27		0x108a1ce		488d0d90890100		LEAQ go.string.*+4133(SB), CX
	  main.go:27		0x108a1d5		48890c24		MOVQ CX, 0(SP)
	  main.go:27		0x108a1d9		48c74424080b000000	MOVQ $0xb, 0x8(SP)
	  main.go:28		0x108a1e2		488d0c24		LEAQ 0(SP), CX
	  main.go:28		0x108a1e6		488708			XCHGQ CX, 0(AX)
	  main.go:29		0x108a1e9		488b6c2410		MOVQ 0x10(SP), BP
	  main.go:29		0x108a1ee		4883c418		ADDQ $0x18, SP
	  main.go:29		0x108a1f2		c3			RET
*/
func main() {
	u8 := NewUint64()
	u8.Store()
	b := u8.Load()
	fmt.Println(len(b))
	b2 := u8.Load()
	fmt.Println(len(b2))
}


