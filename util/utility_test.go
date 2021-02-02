package util

import (
	"go-learn/core"
	"math/rand"
	"testing"
)

func BenchmarkIsOddBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsOddBinary(rand.Int31())
	}
}

func BenchmarkIsOddMod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsOddMod(rand.Int31())
	}
}

func BenchmarkSumSeq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var w core.Worker = &core.CustomWorker{}
		input := w.GenerateIntegerSlice(3000000)
		SumSeq(input)
	}
}
