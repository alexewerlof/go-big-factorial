package main

import (
	"math/big"
)

func toBig(x uint64) *big.Int {
	return big.NewInt(int64(x))
}
