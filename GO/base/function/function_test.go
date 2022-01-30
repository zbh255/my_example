package function

import "testing"

// 匿名函数的一些行为
// 代码思路来自《GO语言学习笔记》
func test() []func() {
	funs := make([]func(), 0, 2)
	for i := 0; i < 2; i++ {
		funs = append(funs, func() {
			println(&i, i)
		})
	}
	return funs
}

func TestFunction(t *testing.T) {
	for _, v := range test() {
		v()
	}
}
