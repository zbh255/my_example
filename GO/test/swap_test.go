package test

import "testing"

func Swap(x,y *int) {
	*y = *x ^ *y
	*x = *x ^ *y
	*y = *x ^ *y
}

func CSwap(x,y *int) {
	n := *x
	*x = *y
	*y = n
}

func MSwap(x,y *int) {
	*y = *x + *y
	*x = *y - *x
}

func BenchmarkSwap(b *testing.B) {
	b.Run("XOR", func(b *testing.B) {
		x,y := 7,9
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Swap(&x,&y)
		}
	})
	b.Run("Center", func(b *testing.B) {
		x,y := 7,9
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			CSwap(&x,&y)
		}
	})
	b.Run("Math", func(b *testing.B) {
		x,y := 7,9
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			MSwap(&x,&y)
		}
	})
}

