package main

import (
	"sync"
	"testing"
	"unsafe"
)

type Interface struct {
	_iface uintptr
	data uintptr
}

// sync.RWMutex返回的Locker接口其实是可以被还原的
func TestUnsafeConversion(t *testing.T) {
	lock := sync.RWMutex{}
	lock.Lock()
	readLock := lock.RLocker()
	readWriteLock := (*sync.RWMutex)(unsafe.Pointer((*Interface)(unsafe.Pointer(&readLock)).data))
	readWriteLock.Unlock()
}

// 安全的类型断言转换失败
func TestSafeConversion(t *testing.T) {
	lock := sync.RWMutex{}
	lock.Lock()
	readLock := lock.RLocker()
	readWriteLock := readLock.(*sync.RWMutex)
	readWriteLock.Unlock()
}