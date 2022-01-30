package base

import "testing"

func TestArrayAddressable(t *testing.T) {
	//var array = map[string][2]int{
	//	"abc":{1,2},
	//}

	//slice := array["abc"][:]
}

// 测试逃逸到堆的切片和数组复制访问性能的差别
// 对于小数组来说拷贝整个数组，未必要比逃逸到堆的切片访问性能要低
// 代码思路来自《GO语言学习笔记》
// -gcflags '-m'
func BenchmarkEscapes(b *testing.B) {
	b.Run("SliceEscapes", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			slice()
		}
	})
	b.Run("ArrayNoEscapes", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			array()
		}
	})
}

func slice() []int {
	x := make([]int, 1024)
	for k := range x {
		x[k] = k
	}
	return x
}

func array() [1024]int {
	var x [1024]int
	for k := range x {
		x[k] = k
	}
	return x
}
