package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func listenSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1,
		syscall.SIGUSR2, syscall.SIGTSTP)
	select {
	case <-sigs:
		fmt.Println("exitapp,sigs:", sigs)
		os.Exit(0)
	}
}

func main() {
	fmt.Println("start...")
	go listenSignal()
	sum := 0
	for {
		sum++
		fmt.Println(sum)
		time.Sleep(time.Second)
	}
}
