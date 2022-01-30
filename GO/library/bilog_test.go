package library

import (
	"github.com/golang/glog"
	"github.com/zbh255/bilog"
	"testing"
)

type Write struct {

}

func (w *Write) Write(p []byte) (int, error) {
	return len(p),nil
}

func BenchmarkLogger(b *testing.B) {
	b.Run("Bilog", func(b *testing.B) {
		b.ReportAllocs()
		logger := bilog.NewLogger(&Write{},bilog.PANIC)
		for i := 0; i < b.N; i++ {
			logger.Info("hello world")
			logger.Trace("hello world")
			logger.Flush()
		}
	})
	b.Run("Glog", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			glog.Info("hello world")
			glog.Error("hello world")
		}
	})
}
