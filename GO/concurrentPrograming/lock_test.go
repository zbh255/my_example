package main

import (
	"sync"
	"testing"
)

// 测试读写锁和普通互斥锁在面对读多写少的情况下的性能
func BenchmarkMutexAndRWMutex(b *testing.B) {
	b.Run("Mutex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Mutex(100000)
		}
	})
	b.Run("RWMutex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			RWMutex(100000)
		}
	})
}

func Mutex(n int) {
	wg := sync.WaitGroup{}
	wg.Add(n + 2)
	count := &MutexCont{}
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < n/10; i++ {
				count.Count()
				wg.Done()
			}
		}()
	}
	for i := 0; i < 2; i++ {
		go func() {
			count.CountAdd()
			wg.Done()
		}()
	}
	wg.Wait()
}

func RWMutex(n int) {
	wg := sync.WaitGroup{}
	wg.Add(n + 2)
	count := &RWMutexCont{}
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < n/10; i++ {
				count.Count()
				wg.Done()
			}
		}()
	}
	for i := 0; i < 2; i++ {
		go func() {
			count.CountAdd()
			wg.Done()
		}()
	}
	wg.Wait()
}

type MutexCont struct {
	mu    sync.Mutex
	count int
}

func (m *MutexCont) Count() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.count
}

func (m *MutexCont) CountAdd() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.count++
}

type RWMutexCont struct {
	mu    sync.RWMutex
	count int
}

func (m *RWMutexCont) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.count
}

func (m *RWMutexCont) CountAdd() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.count++
}
