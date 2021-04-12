package main

import (
	"fmt"
	"math/big"
)

const log = false

func pow(a *big.Int, b uint64, result chan<- *big.Int) {
	var res big.Int
	if log {
		fmt.Println(a, "^", b)
	}
	result <- res.Exp(a, toBig(b), nil)
}

func mul(values chan *big.Int, done chan<- bool) {
	var res big.Int
	if log {
		fmt.Println("x")
	}
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
