package main

import (
	"fmt"
	"math/big"
	"time"
)

var zero = big.NewInt(0)
var one = big.NewInt(1)
var two = big.NewInt(2)

type PowArgs struct {
	x *big.Int
	n uint64
}

func pow(ppBig PowArgs) *big.Int {
	if log {
		// fmt.Println(a, "^", b)
		var start = time.Now()
		defer fmt.Println("^", time.Since(start))
	}
	var res big.Int
	if ppBig.x.Cmp(one) == 0 {
		return ppBig.x
	} else if ppBig.x.Cmp(two) == 0 {
		// It is much faster to compute powers of two by just setting a big on an all-zero number
		return res.SetBit(zero, int(ppBig.n), 1)
	} else {
		return res.Exp(ppBig.x, toBig(ppBig.n), nil)
	}
}

type MulArgs struct {
	a *big.Int
	b *big.Int
}

func mul(m MulArgs) *big.Int {
	if log {
		// fmt.Println(a, "×", b)
		var start = time.Now()
		defer fmt.Println("×", time.Since(start))
	}
	var res big.Int
	return res.Mul(m.a, m.b)
}
