package base

import (
	"fmt"
	"testing"
)

// TestRange 用于测试Go range 的一些行为
func TestRange(t *testing.T) {
	value := []int{1, 2, 3, 4}
	copyValue := make([]*int, 0, len(value))
	for _, v := range value {
		copyValue = append(copyValue, &v)
	}
	for _, v := range copyValue {
		fmt.Println(*v)
	}
}
