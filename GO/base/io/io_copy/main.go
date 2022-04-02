package main

import (
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		panic(err)
	}
	io.Copy(conn,nil)
}