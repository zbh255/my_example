//go:build darwin
// +build darwin

package main

import "golang.org/x/sys/unix"

func main() {
	unix.Kqueue()
}