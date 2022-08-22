package benchmarks

import (
	"github.com/GalloaFranco/golang-katas/katas"
	"testing"
)

func BenchmarkGetMiddleBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		katas.GetMiddleBytes("Middle")
	}
}

func BenchmarkGetMiddleStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		katas.GetMiddleStrings("Middle")
	}
}
