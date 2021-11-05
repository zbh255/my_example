package main

import (
	"os"
	"os/exec"
)

func main() {

}

func background(logPath string) {
	_ = exec.Command(os.Args[0],os.Args[1:]...)
	
}