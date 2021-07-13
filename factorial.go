package main

import (
	"fmt"
	"math/big"
	"runtime"
	"time"
)

const log = false

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

func powWorker(powChan <-chan PowArgs, feedChan chan<- *big.Int) {
	for ppBig := range powChan {
		feedChan <- pow(ppBig)
		if log {
			fmt.Println("feedChan <- ", ppBig.x.String(), " ** ", ppBig.n)
		}
	}
}

func feedMulWorkers(feedChan <-chan *big.Int, mulChan chan<- MulArgs, times int) {
	for i := 0; i < times; i++ {
		mulChan <- MulArgs{a: <-feedChan, b: <-feedChan}
		if log {
			fmt.Println("mulChan <- two numbers")
		}
	}
}

func mulWorker(mulChan <-chan MulArgs, feedChan chan<- *big.Int) {
	for mulBig := range mulChan {
		feedChan <- mul(mulBig)
		if log {
			fmt.Println("feedChan <- ", mulBig.a.String(), " x ", mulBig.b.String())
		}
	}
}

func factorial(n uint64) *big.Int {
	fmt.Println("Digesting...")
	primePowers := allMergedReducedPrimeFactors(n)
	powsLen := len(primePowers)
	//fmt.Println("Digested to", pows)

	numWorkers := runtime.GOMAXPROCS(0) / 2
	powChan := make(chan PowArgs, powsLen)
	feedChan := make(chan *big.Int, powsLen)
	mulChan := make(chan MulArgs, powsLen)

	fmt.Println("Using", numWorkers, "workers for powering and multiplication")
	for i := 0; i < numWorkers; i++ {
		go powWorker(powChan, feedChan)
		go mulWorker(mulChan, feedChan)
	}

	for n, x := range primePowers {
		if log {
			fmt.Println("powChan <- ", x.String(), n)
		}
		powChan <- PowArgs{x, n}
	}

	feedMulWorkers(feedChan, mulChan, powsLen-1)

	primePowers = nil

	defer fmt.Println("\nDone!")
	return <-feedChan
}
