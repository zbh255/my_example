package benchmark

import "testing"

/*
	测试Xor和Mov的清零方法的性能
	关闭编译器优化及内联: -gcflags "-N -l"
*/
//goos: darwin
//goarch: amd64
//pkg: example/benchmark
//cpu: Intel(R) Core(TM) i7-8705G CPU @ 3.10GHz
//BenchmarkXorAndMov
//BenchmarkXorAndMov/Xor
//BenchmarkXorAndMov/Xor-8         	543602060	         2.163 ns/op	       0 B/op	       0 allocs/op
//BenchmarkXorAndMov/Mov
//BenchmarkXorAndMov/Mov-8         	538426809	         2.186 ns/op	       0 B/op	       0 allocs/op
func BenchmarkXorAndMov(b *testing.B) {
	b.Run("Xor", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			b := 0xf67556
			b = b ^ b
		}
	})
	b.Run("Mov", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = 0xf67556
			_ = 0
		}
	})
}
