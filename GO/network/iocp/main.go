//go:build windows
// +build windows
package main

import (
	"golang.org/x/sys/windows"
)

func main() {
	windows.CreateIoCompletionPort()
}