package main

import (
	"math"
	"math/big"
)

func toBig(x uint64) *big.Int {
	return big.NewInt(int64(x))
}

func sqrt(x uint64) uint64 {
	return uint64(math.Sqrt(float64(x)))
}
