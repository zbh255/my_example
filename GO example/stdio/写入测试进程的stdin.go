package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	cmd := exec.Command("./read")
	cmd.Stdout = os.Stdout
	cmd.Stdin = strings.NewReader(`{"name":"hello","test":"world"}`)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 2)
}
