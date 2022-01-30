package pointer

import (
	"testing"
	"unsafe"
)

// 这里研究Go指针与空接口的一些行为，包括多重指针等

type MultiPtr struct {
	_type *MultiPtr
	data  uint64
}

// 测试多重指针
func TestStructMultiPtr(t *testing.T) {
	multis := &MultiPtr{
		&MultiPtr{
			_type: &MultiPtr{
				_type: &MultiPtr{
					_type: nil,
					data:  4,
				},
				data: 3,
			},
			data: 2,
		}, 1,
	}
	// multis的内存布局
	// ****MultiPtr
	itab := ***(***uintptr)(unsafe.Pointer(multis))
	mItab := ****(****MultiPtr)(unsafe.Pointer(multis))
	multi := (*MultiPtr)(unsafe.Pointer(itab))
	t.Log(multi.data)
	t.Log(mItab.data)
}

// 测试多重指针与空接口
func TestPointer(t *testing.T) {
	user := struct {
		name string
		id   string
	}{
		"zbh255",
		"1234567890",
	}
	userI := interface{}(user)

	itab1 := (uintptr)(unsafe.Pointer(&userI))
	itab2 := (*uintptr)(unsafe.Pointer(&userI))
	itab3 := (**uintptr)(unsafe.Pointer(&userI))
	itab4 := *(**uintptr)(unsafe.Pointer(&userI))
	t.Log(itab1)
	t.Log(itab2)
	t.Log(itab3)
	t.Log(itab4)

	itab := *(***eface)(unsafe.Pointer(&userI))
	t.Log(unsafe.Pointer(itab))
}
