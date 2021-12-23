package main

import (
	"os"
	"os/exec"
)

func main()  {
	cmd := exec.Command("echo","shell")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
