package base

import (
	"sync"
	"testing"
)

// go version 1.17

func call() {
	var mu sync.Mutex
	mu.Lock()
	mu.Unlock()
}

func deferCall() {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
}

/*
	goos: darwin
	goarch: amd64
	pkg: example/base
	cpu: Intel(R) Core(TM) i7-8705G CPU @ 3.10GHz
	BenchmarkCall
	BenchmarkCall/NoDefer
	BenchmarkCall/NoDefer-8         	52188236	        22.49 ns/op
	BenchmarkCall/Defer
	BenchmarkCall/Defer-8           	50228857	        24.85 ns/op
	PASS
*/
func BenchmarkCall(b *testing.B) {
	b.Run("NoDefer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			call()
		}
	})
	b.Run("Defer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			deferCall()
		}
	})

}
