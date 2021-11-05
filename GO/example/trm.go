package main

import (
	"strings"
)

func main() {
	println(strings.TrimPrefix("/Users/cache/path1/path2", "/Users/cache/path1"+`/`))
}
