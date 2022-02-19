package benchmark

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// 测试从strconv int转换到string的开销

func BenchmarkConvert(b *testing.B) {
	b.Run("Itoa", func(b *testing.B){
		b.ReportAllocs()
		ints := random1024Int()
		for i := 0; i < b.N; i++ {
			for i := 0; i < len(ints); i++ {
				_ = strconv.Itoa(ints[i])
			}
		}
	})
	b.Run("Hash", func(b *testing.B){
		b.ReportAllocs()
		hashTab := make(map[int]string)
		hashTab[1024] = "1024"
		for i := 0; i < b.N; i++ {
			_ = hashTab[1024]
		}
	})
	b.Run("Cache", func(b *testing.B) {
		b.ReportAllocs()
		hashTab := make(map[int]string,390)
		ints := random1024Int()
		for i := 0; i < b.N; i++ {
			for j := 0; j < len(ints); j++ {
				t,ok := hashTab[ints[j]]
				if !ok {
					hashTab[ints[j]] = strconv.Itoa(ints[j])
					t = hashTab[ints[j]]
				}
				_ = t
			}
		}
	})
}

func random1024Int() [102400]int{
	tmp := [102400]int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(tmp); i++ {
		tmp[i] = rand.Intn(224) + 100
	}
	return tmp
}