//go:build linux
// +build linux

package main

import (
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	var buffer [256]byte
	_, err := conn.Read(buffer[:])
	if err != nil {
		return
	}
	var buf []byte
	buf = append(buf, "HTTP/1.1 200 OK\r\nServer: gnet\r\nContent-Type: text/plain\r\nDate: "...)
	buf = time.Now().AppendFormat(buf, "Mon, 02 Jan 2006 15:04:05 GMT")
	buf = append(buf, "\r\nContent-Length: 12\r\n\r\nHello World!"...)
	_, err = conn.Write(buf)
	if err != nil {
		return
	}
}