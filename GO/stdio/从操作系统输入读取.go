package main

import (
	"os"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)
	tmp := make([]byte, 512)
	_, _ = os.Stdin.Read(tmp)
	_, _ = os.Stdout.Write(tmp)
}
