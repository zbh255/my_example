package main

import "testing"

// 循环一
func TestLoop1(t *testing.T) {
	arr := [64 * 1024 * 1024]int32{}
	for i := 0; i < len(arr); i++ {
		arr[i] *= 3
	}
}

// 循环二
func TestLoop2(t *testing.T) {
	arr := [64 * 1024 * 1024]int32{}
	for i := 0; i < len(arr); i += 16 {
		arr[i] *= 3
	}
}

// 二维数组循环1
// 按行循环
func TestLoop3(t *testing.T) {
	arr := [64 * 1024 * 1024][64]int32{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			arr[i][j] *= 3
		}
	}
}

// 二维数组循环2
// 按列循环
func TestLoop4(t *testing.T) {
	arr := [64 * 1024 * 1024][64]int32{}
	for i := 0; i < len(arr[i]); i++ {
		for j := 0; j < len(arr); j++ {
			arr[j][i] *= 3
		}
	}
}
