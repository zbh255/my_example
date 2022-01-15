package escapes

import (
	"testing"
	_ "unsafe"
)

// 测试Go 逃逸分析的一些行为

type User struct {
	id   string
	name string
}

func Heap(n int) {
	user := NewUser()
	for i := 0; i < n; i++ {
		_ = user.id
		_ = user.name
	}
}

func Stack(n int) {
	user := &User{
		id:   "xxl",
		name: "zbh255",
	}
	for i := 0; i < n; i++ {
		_ = user.id
		_ = user.name
	}
}

func NewUser() *User {
	return &User{
		id:   "xxl",
		name: "zbh255",
	}
}

// 测试变量被分配到栈和堆具体访问性能区别
// 使用参数: -gcflags '-m -l'
func BenchmarkHeapAndStack(b *testing.B) {
	b.Run("EscapesToHeap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Heap(10000)
		}
	})

	b.Run("NoEscapes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Stack(10000)
		}
	})
}

type BigObjUser struct {
	id   string
	name string
	data []byte
}

func NewBigObjUser() *BigObjUser {
	return &BigObjUser{
		id:   "xxl",
		name: "zbh255",
		data: make([]byte, 1024*1024),
	}
}

func BigObjHeap(n int) {
	user := NewBigObjUser()
	for i := 0; i < n; i++ {
		_ = user.id
		_ = user.name
		_ = user.data
	}
}

func BigObjStack(n int) {
	user := &BigObjUser{
		id:   "xxl",
		name: "zbh255",
		data: make([]byte, 1024*1024),
	}
	for i := 0; i < n; i++ {
		_ = user.id
		_ = user.name
		_ = user.data
	}
}

// 测试分配大对象的访问性能
// 参数: -gcflags '-m -l'
func BenchmarkBigObjToHeapAndStack(b *testing.B) {
	b.Run("BigObjEscapesToHeap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BigObjHeap(10000)
		}
	})
	b.ResetTimer()
	b.Run("BigObjNoEscapes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BigObjStack(10000)
		}
	})
}
