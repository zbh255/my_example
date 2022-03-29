package network

import (
	"net"
	"sync"
	"testing"
)

func BenchmarkMultiClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		dial(2000,"192.168.1.228:9080")
	}
}

func dial(n int,addr string) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				panic(err)
			}
			var buf [128]byte
			_, err = conn.Write([]byte("hello world!"))
			if err != nil {
				panic(err)
			}
			_, err = conn.Read(buf[:])
			if err != nil {
				panic(err)
			}
		}()
	}
	wg.Wait()
}