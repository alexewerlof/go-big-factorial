package main

import (
	"fmt"
	"math/big"
	"time"
)

const log = false

var zero = big.NewInt(0)
var one = big.NewInt(1)
var two = big.NewInt(2)

func pow(a *big.Int, b uint64, result chan<- *big.Int) {
	if log {
		// fmt.Println(a, "^", b)
		var start = time.Now()
		defer fmt.Println("^", time.Since(start))
	}
	var res big.Int
	if a.Cmp(one) == 0 {
		result <- a
	} else if a.Cmp(two) == 0 {
		// It is much faster to compute powers of two by just setting a big on an all-zero number
		result <- res.SetBit(zero, int(b), 1)
	} else {
		result <- res.Exp(a, toBig(b), nil)
	}
}

func mul(values chan *big.Int, done chan<- bool) {
	if log {
		// fmt.Println(a, "×", b)
		var start = time.Now()
		defer fmt.Println("×", time.Since(start))
	}
	var res big.Int
	values <- res.Mul(<-values, <-values)
	done <- true
}

func factorial(n uint64) *big.Int {
	fmt.Println("Digesting...")
	primePowers := allMergedReducedPrimeFactors(n)
	powsLen := len(primePowers)
	vals := make(chan *big.Int, powsLen)
	//fmt.Println("Digested to", pows)
	fmt.Println("\nPowering...")
	for power, prime := range primePowers {
		if log {
			fmt.Println(prime, " ** ", power)
		}

		if power == 1 {
			vals <- prime
		} else {
			go pow(prime, power, vals)
		}
	}
	primePowers = nil
	fmt.Println("\nMultiplying...")
	multiOp := make(chan bool)
	for i := 1; i < powsLen; i++ {
		go mul(vals, multiOp)
	}
	for i := 1; i < powsLen; i++ {
		x := <-multiOp
		if log {
			fmt.Println("One multiplication operation was done", i, x)
		}
	}
	fmt.Println("\nDone!")
	return <-vals
}
