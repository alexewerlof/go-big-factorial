package main

import (
	"fmt"
	"math/big"
	"time"
)

var zero = big.NewInt(0)
var one = big.NewInt(1)
var two = big.NewInt(2)

func pow(x *big.Int, n uint64) *big.Int {
	if log {
		// fmt.Println(a, "^", b)
		var start = time.Now()
		defer fmt.Println("^", time.Since(start))
	}
	var res big.Int
	if x.Cmp(one) == 0 {
		return x
	} else if x.Cmp(two) == 0 {
		// It is much faster to compute powers of two by just setting a big on an all-zero number
		return res.SetBit(zero, int(n), 1)
	} else {
		return res.Exp(x, toBig(n), nil)
	}
}

func mul(a, b *big.Int) *big.Int {
	if log {
		// fmt.Println(a, "×", b)
		var start = time.Now()
		defer fmt.Println("×", time.Since(start))
	}
	var res big.Int
	return res.Mul(a, b)
}
