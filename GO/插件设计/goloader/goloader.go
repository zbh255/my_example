package main

import (
	"fmt"
	"github.com/pkujhd/goloader"
	"unsafe"
)

func main()  {
	symPtr := map[string]uintptr{}
	err := goloader.RegSymbol(symPtr)
	if err != nil {
		panic(err)
	}
	goloader.RegTypes(symPtr,fmt.Fprintln)
	linker,err := goloader.ReadObjs([]string{"./upload.o"},[]string{""})
	if err != nil {
		panic(err)
	}
	codeModule, err := goloader.Load(linker, symPtr)
	if err != nil {
		panic(err)
	}
	// select function
	iFn := codeModule.Syms["main.Print"]
	if iFn == 0 {
		panic("the function is not found")
	}
	funcPtrContainer := (uintptr)(unsafe.Pointer(&iFn))
	Print := *(*func())(unsafe.Pointer(&funcPtrContainer))
	Print()
}