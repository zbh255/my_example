package benchmark

import (
	"net"
	"testing"
)

// 编译参数: -gcflags "-N -l"
// @Feature: 关闭编译器优化和内联来测试函数调用的耗时
func BenchmarkFunctionCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hello(nil)
	}
}

func hello(conn net.Conn) {
	_ = conn
}
