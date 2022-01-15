package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSharingMemory(t *testing.T) {
	var x int
	var done int

	go func() {
		for done == 0 {
			fmt.Printf("done = %d\n", done)
		}
		fmt.Printf("x = %d\n", x)
	}()

	go func() {
		x = 1
		done = 1
	}()

	for {
		time.Sleep(1)
	}
}
