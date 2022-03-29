package main

import (
	"net"
	"testing"
)

func TestEpollLink(t *testing.T) {
	conn, err := net.Dial("tcp", "192.168.1.150:8090")
	if err != nil {
		t.Error(err)
		return
	}
	conn.Write([]byte("hello world!"))
	var buf [512]byte
	conn.Read(buf[:])
	conn.Close()
}
