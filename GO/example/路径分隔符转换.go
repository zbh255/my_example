package main

import (
	"path/filepath"
)

func main() {
	println(filepath.FromSlash("/conf/dev/app.conf/"))
	println(filepath.FromSlash("\\conf\\dev\\app.conf"))
	println(filepath.ToSlash("/conf/dev/app.conf"))
}
