package typeSystem

import (
	"fmt"
	"testing"
)

// 代码思路来自《GO语言学习笔记》
// struct 的标签也是类型的组成部分
func TestTyeSystemToStruct(t *testing.T) {
	var a = struct {
		x int `toml:"x"`
		d int `toml:"d"`
	}{}
	var b = struct {
		x int
		d int
	}{}
	// 类型不匹配
	//t.Log(a == b)
	t.Log(a)
	t.Log(b)

}

type N int

func (n N) String() {
	fmt.Printf("ptr: %p,value: %d\n",&n,n)
}

// 代码思路来自《GO语言学习笔记》
func TestMethodExpressions(t *testing.T) {
	var n N = 90

	// 被注释的代码在go 1.17不能正确地运行
	//test := n.String
	//test(n)
	N.String(n)
	(*N).String(&n)
}