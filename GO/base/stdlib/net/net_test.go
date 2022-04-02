package main

import (
	"net"
	"sync"
	"testing"
)

func TestTcpFd(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()
			conn, err := net.Dial("tcp", "127.0.0.1:9090")
			if err != nil {
				t.Error(err)
				return
			}
			conn.Close()
		}()
	}
	wg.Wait()
}

//	@Feature: 测试使用反射的方式获取SysFd和不使用反射获取的性能
//	goos: darwin
//	goarch: amd64
//	pkg: example/base/stdlib/net
//	cpu: Intel(R) Core(TM) i7-8705G CPU @ 3.10GHz
//	BenchmarkGetFdPreference
//	BenchmarkGetFdPreference/Reflect
//	BenchmarkGetFdPreference/Reflect-8         	 4006008	       270.9 ns/op	      32 B/op	       4 allocs/op
//	BenchmarkGetFdPreference/NoReflect
//	BenchmarkGetFdPreference/NoReflect-8       	1000000000	         0.3713 ns/op	       0 B/op	       0 allocs/op
//	PASS
func BenchmarkGetFdPreference(b *testing.B) {
	listener, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		b.Error(err)
		return
	}
	go func() {
		dial, err := net.Dial("tcp", "127.0.0.1:9090")
		if err != nil {
			b.Error(err)
			return
		}
		_, _ = dial.Write([]byte("hello world!"))
		dial.Close()
	}()
	conn, err := listener.Accept()
	if err != nil {
		b.Error(err)
		return
	}
	b.Run("Reflect", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = SocketFD(conn)
		}
	})
	b.Run("NoReflect", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = Fd(conn)
		}
	})
}
