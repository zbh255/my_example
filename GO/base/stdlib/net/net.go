package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"reflect"
	"unsafe"
)

// 验证标准库net包的一些行为

func main() {
	listener,err := net.Listen("tcp","127.0.0.1:9090")
	if err != nil {
		panic(err)
	}

	for {
		conn,err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			fmt.Println(Fd(conn))
			conn.Close()
		}()
	}
}

func Fd(conn net.Conn) int32 {
	pollFdLdByte := *(**[24]byte)((*[2]unsafe.Pointer)(unsafe.Pointer(&conn))[1])
	return int32(binary.LittleEndian.Uint64(pollFdLdByte[16:]))
}

// SocketFD from https://colobu.com/2019/02/23/1m-go-tcp-connection/
func SocketFD(conn net.Conn) int {
	//tls := reflect.TypeOf(conn.UnderlyingConn()) == reflect.TypeOf(&tls.Conn{})
	// Extract the file descriptor associated with the connection
	//connVal := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn").Elem()
	tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	//if tls {
	//	tcpConn = reflect.Indirect(tcpConn.Elem())
	//}
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")
	return int(pfdVal.FieldByName("Sysfd").Int())
}