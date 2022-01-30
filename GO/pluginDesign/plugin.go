package main

import (
	"os"
	"plugin"
)

func main() {
	p,_ := os.Getwd()
	pg,err := plugin.Open(p + "/log_plugin")
	if err != nil {
		panic(err)
	}
	os.Args = []string{"1","2","我是参数列表"}
	sb,_ := pg.Lookup("Print")
	sb.(func(string))("hw")
}