package test

import (
	"math"
	"testing"
	"unsafe"
)

// 卡马克算法老代码
//func Qrsqrt(number float32) float32 {
//	var i int32
//	var x2,y float32
//	const threeHalfs float32 = 1.5;
//
//	x2 = number * 0.5
//	y = number
//	i = *(*int32)(unsafe.Pointer(&y))
//	i = 0x5f3759df - (i >> 1)
//	y = *(*float32)(unsafe.Pointer(&i))
//	y = y * (threeHalfs - (x2 * y * y))
//	return y * number
//}

func Qrsqrt(number float32) float32 {
	y := number
	x2 := number * 0.5
	var halfs float32 = 1.5
	i := *(*uint32)(unsafe.Pointer(&y))
	i = 0x5f3759df - (i >> 1)
	y = *(*float32)(unsafe.Pointer(&i))
	y = y * (halfs - (x2 * y * y))
	return y * number
}

func Crsqrt(number float64) float64 {
	i := number / 2
	a := 0.0000000000001
	count := 0
	for math.Abs(math.Pow(i, 2)-number) > a {
		i = i - (math.Pow(i, 2)-number)/(2*i)
		count++
	}
	return i
}

func InvSqrt(x float32) float32 {
	var xhalf float32 = 0.5 * x // get bits for floating VALUE
	i := math.Float32bits(x)    // gives initial guess y0
	i = 0x5f375a86 - (i >> 1)   // convert bits BACK to float
	x = math.Float32frombits(i) // Newton step, repeating increases accuracy
	x = x * (1.5 - xhalf*x*x)
	x = x * (1.5 - xhalf*x*x)
	x = x * (1.5 - xhalf*x*x)
	return 1 / x
}

//func BenchmarkSqrt(b *testing.B) {
//	b.Run("Std", func(b *testing.B) {
//		b.ResetTimer()
//		for i := 0; i < b.N; i++ {
//			math.Sqrt(500)
//		}
//	})
//	b.Run("Custom", func(b *testing.B) {
//		b.ResetTimer()
//		for i := 0; i < b.N; i++ {
//			Qrsqrt(500)
//		}
//	})
//	b.Run("Custom3", func(b *testing.B) {
//		b.ResetTimer()
//		for i := 0; i < b.N; i++ {
//			Qrsqrt2(500)
//		}
//	})
//	b.Run("Inv", func(b *testing.B) {
//		b.ResetTimer()
//		for i := 0; i < b.N; i++ {
//			InvSqrt(500)
//		}
//	})
//	b.Run("Custom2", func(b *testing.B) {
//		b.ResetTimer()
//		for i := 0; i < b.N; i++ {
//			Crsqrt(500)
//		}
//	})
//}

func BenchmarkSqrt(b *testing.B) {
	b.Run("Std Lib", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			math.Sqrt(500)
		}
	})
	b.Run("CarMock", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Qrsqrt(500)
		}
	})
	b.Run("NewtonIter", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Crsqrt(500)
		}
	})
}

func TestSqrt(t *testing.T) {
	t.Log(Crsqrt(500))
	t.Log(Qrsqrt(4))
	t.Log(InvSqrt(500))
	t.Log(math.Sqrt(500))
}
