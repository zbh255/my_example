package escapes

// 测试atomic和unsafe下逃逸分析的正确性

import (
	"math"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

type TestFactory struct {
	i1 int32
	buf *[32]byte
}

func NewTestFactory() *TestFactory {
	return &TestFactory{
		buf: &[32]byte{},
	}
}

func random32Bytes() [32]byte {
	rand.Seed(time.Now().UnixNano())
	var tmp [32]byte
	for i := range tmp {
		tmp[i] = byte(rand.Intn(math.MaxUint8))
	}
	return tmp
}

func (f *TestFactory) Store() {
	tmp := random32Bytes()
	atomic.StoreUintptr((*uintptr)(unsafe.Pointer(f.buf)),(uintptr)(unsafe.Pointer(&tmp)))
}

func (f *TestFactory) Load() [32]byte {
	bufPtr := atomic.LoadUintptr((*uintptr)(unsafe.Pointer(f.buf)))
	return *(*[32]byte)(unsafe.Pointer(bufPtr))
}

func TestUnsafeEscapes(t *testing.T) {
	test := NewTestFactory()
	test.Store()
	t.Log(test.Load())
}