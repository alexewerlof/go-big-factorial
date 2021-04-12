package main

import (
	"testing"
)

func BenchmarkFib10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		factorial(uint64(b.N * 1000))
	}
}
