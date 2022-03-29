package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	pprof.StartCPUProfile(f)
	log.Println()
	defer pprof.StopCPUProfile()
	for i := 0; i < 1000000;i++ {
		_,_,_, _ = runtime.Caller(2)
	}
}