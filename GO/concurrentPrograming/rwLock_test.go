package main

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Interface struct {
	_iface uintptr
	data uintptr
}

// sync.RWMutex返回的Locker接口其实是可以被还原的
func TestRwLock(t *testing.T) {
	lock := sync.RWMutex{}
	lock.Lock()
	refLock := lock.RLocker()
	of := unsafe.Sizeof(refLock)
	fmt.Println(of)
	rwxInf := *(*Interface)(unsafe.Pointer(&refLock))
	rwx := (*sync.RWMutex)(unsafe.Pointer(rwxInf.data))
	rwx.Unlock()
	rwx.Lock()
	rwx.Unlock()
}
