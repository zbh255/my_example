package typeAssertion

import (
	"fmt"
	"testing"
	"unsafe"
)

type Interface interface {

}

type impl struct {
}

// 内存不可读
func TestIface(t *testing.T) {
	var inter Interface = &impl{}
	iface := (*iface)(unsafe.Pointer(&inter))
	fmt.Println(iface)
	fmt.Println(iface.tab.inter)
}
