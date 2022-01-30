package main

import (
	"sync/atomic"
	"testing"
	"unsafe"
)

// 测试atomic.Value和LoadUintptr之间的性能差距

type Value struct {
	val atomic.Value
}

func (v *Value) Load() [32]byte {
	return *(v.val.Load().(*[32]byte))
}

func (v *Value) Store() {
	v.val.Store(&[32]byte{65,91})
}

type Uintptr struct {
	buf *[32]byte
}

func (u *Uintptr) Load() [32]byte {
	return *(*[32]byte)(unsafe.Pointer(atomic.LoadUintptr((*uintptr)(unsafe.Pointer(&u.buf)))))
}

func (u *Uintptr) Store() {
	atomic.StoreUintptr((*uintptr)(unsafe.Pointer(&u.buf)),(uintptr)(unsafe.Pointer(&[32]byte{56,91})))
}

func BenchmarkAtomic(b *testing.B) {
	b.Run("ValueLoad", func(b *testing.B) {
		b.ReportAllocs()
		val := &Value{}
		val.Store()
		for i := 0; i < b.N; i++ {
			_ = val.Load()
		}
	})
	b.Run("ValueStore", func(b *testing.B) {
		b.ReportAllocs()
		val := &Value{}
		for i := 0; i < b.N; i++ {
			val.Store()
		}
	})
	b.Run("UintptrLoad", func(b *testing.B) {
		b.ReportAllocs()
		u := &Uintptr{}
		u.Store()
		for i := 0; i < b.N; i++ {
			_ = u.Load()
		}
	})
	b.Run("UintptrStore", func(b *testing.B) {
		b.ReportAllocs()
		u := &Uintptr{}
		for i := 0; i < b.N; i++ {
			u.Store()
		}
	})
}

