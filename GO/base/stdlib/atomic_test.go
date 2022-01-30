package main

import (
	"fmt"
	"reflect"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

type Test struct {
	i1 int64
	i2 byte
	i3 int64
}

// 编译32位程序, GOARCH=386
func TestAtomicAlign32Bit(t *testing.T) {
	t1 := Test{}
	t.Log(atomic.AddInt64(&t1.i3,100))
	t.Log(runtime.GOARCH)
}

func NewNoUintPtr() *NoUintPtr {
	return &NoUintPtr{}
}

func NewUintPtr() *UintPtr {
	return &UintPtr{}
}

func NewPointer() *Pointer {
	return &Pointer{}
}

func NewUint64() *Uint64 {
	return &Uint64{}
}

type NoUintPtr struct {
	buf *[]byte
}

func (n *NoUintPtr) Load() []byte {
	return *(*[]byte)(unsafe.Pointer(atomic.LoadUintptr((*uintptr)(unsafe.Pointer(&n.buf)))))
}

func (n *NoUintPtr) Store() {
	n.appendBuf()
}

func (n *NoUintPtr) appendBuf() {
	tmp := call()
	atomic.StoreUintptr((*uintptr)(unsafe.Pointer(&n.buf)),(uintptr)(unsafe.Pointer(&tmp)))
}

func call() []byte {
	tmp := make([]byte,0,64)
	tmp = append(tmp,time.Now().String()...)
	return tmp
}

type UintPtr struct {
	buf uintptr
}

func (n *UintPtr) Load() []byte {
	return *(*[]byte)(unsafe.Pointer(atomic.LoadUintptr(&n.buf)))
}

func (n *UintPtr) Store() {
	n.appendBuf()
}

func (n *UintPtr) appendBuf() {
	tmp := call()
	atomic.StoreUintptr(&n.buf,(uintptr)(unsafe.Pointer(&tmp)))
}

type Pointer struct {
	buf unsafe.Pointer
}

func (p *Pointer) Load() []byte {
	return *(*[]byte)(atomic.LoadPointer(&p.buf))
}

func (p *Pointer) Store() {
	p.appendBuf()
}

func (p *Pointer) appendBuf() {
	tmp := call()
	atomic.StorePointer(&p.buf,unsafe.Pointer(&tmp))
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
	tmp := call()
	atomic.StoreUint64(&n.buf,(uint64)(uintptr(unsafe.Pointer(&tmp))))
}

// 测试atomic.XXUintPtr的一些行为
func TestAtomicLoadNoUintPtr(t *testing.T) {
	noPtr := NewNoUintPtr()
	noPtr.Store()
	b := noPtr.Load()
	t.Log(string(b))
	noPtr.Load()
}

func TestAtomicLoadUintPtr(t *testing.T) {
	noPtr := NewUintPtr()
	noPtr.Store()
	b := noPtr.Load()
	fmt.Println(*(*reflect.SliceHeader)(unsafe.Pointer(noPtr.buf)))
	t.Log(string(b))
	fmt.Println(*(*reflect.SliceHeader)(unsafe.Pointer(noPtr.buf)))
	bb := noPtr.Load()
	t.Log(string(bb))
	fmt.Println(*(*reflect.SliceHeader)(unsafe.Pointer(noPtr.buf)))
}

func TestAtomicLoadPointer(t *testing.T) {
	ptr := NewPointer()
	ptr.Store()
	b := ptr.Load()
	t.Log(string(b))
	b2 := ptr.Load()
	t.Log(string(b2))
}

func TestAtomicLoadUint64(t *testing.T) {
	u8 := NewUint64()
	u8.Store()
	b := u8.Load()
	t.Log(len(b))
	bb := u8.Load()
	t.Log(string(bb))
}

func BenchmarkTmp(b *testing.B) {
	b.Run("Tmp", func(b *testing.B){
		b.ReportAllocs()
		factory := NewNoUintPtr()
		factory.Store()
		go func() {
			time.Sleep(time.Millisecond)
			factory.Store()
		}()
		for i := 0; i < b.N; i++ {
			_ = factory.Load()
		}
	})
}