package main

import (
	"math/big"
	"testing"
)

const p = 1000

func BenchmarkSeparate(b *testing.B) {
	pow := big.NewInt(p)
	for i := 0; i < b.N; i++ {
		a := big.NewInt(7)
		var aToPowered big.Int
		aToPowered.Exp(a, pow, nil)
		b := big.NewInt(5)
		var bToPowered big.Int
		bToPowered.Exp(b, pow, nil)
		var res big.Int
		res.Mul(&aToPowered, &bToPowered)
	}
}

func BenchmarkMerged(b *testing.B) {
	pow := big.NewInt(p)
	for i := 0; i < b.N; i++ {
		a := big.NewInt(7)
		b := big.NewInt(5)
		var aByFive big.Int
		aByFive.Mul(a, b)

		var res big.Int
		res.Exp(&aByFive, pow, nil)
	}
}

func Benchmark1(b *testing.B) {
	pow := big.NewInt(1)
	for i := 0; i < b.N; i++ {
		x := big.NewInt(200)
		var res big.Int
		res.Exp(x, pow, nil)
	}
}

func Benchmark1noCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := big.NewInt(200)
		_ = x
	}
}

func Benchmark2(b *testing.B) {
	pow := big.NewInt(p)
	for i := 0; i < b.N; i++ {
		x := big.NewInt(2)
		var res big.Int
		res.Exp(x, pow, nil)
	}
}

func Benchmark2setBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x, res big.Int

		res.SetBit(&x, p, 1)
	}
}

func Benchmark10(b *testing.B) {
	pow := big.NewInt(p)
	for i := 0; i < b.N; i++ {
		x := big.NewInt(10)
		var res big.Int
		res.Exp(x, pow, nil)
	}
}

func BenchmarkNonTen(b *testing.B) {
	pow := big.NewInt(p)
	for i := 0; i < b.N; i++ {
		x := big.NewInt(9)
		var res big.Int
		res.Exp(x, pow, nil)
	}
}

func BenchmarkNonTenMulInPlace(b *testing.B) {
	pow := big.NewInt(p)
	for i := 0; i < b.N; i++ {
		x := big.NewInt(9)
		x.Exp(x, pow, nil)
	}
}
