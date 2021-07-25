package main

import (
	"fmt"
	"math/big"
	"runtime"
)

const log = false

type powArgs struct {
	x *big.Int
	n uint64
}

type mulArgs struct {
	a *big.Int
	b *big.Int
}

func initPowChan(primePowers PriPow, powChan chan<- powArgs) {
	for x, n := range primePowers {
		powChan <- powArgs{toBig(x), n}
	}
	close(powChan)
}

func powWorker(powChan <-chan powArgs, feedChan chan<- *big.Int) {
	for ppBig := range powChan {
		feedChan <- pow(ppBig.x, ppBig.n)
		if log {
			fmt.Println("feedChan <- x^y")
		}
	}
}

func feedChanToMulChan(feedChan <-chan *big.Int, mulChan chan<- mulArgs, times int) *big.Int {
	for i := 0; i < times; i++ {
		mulChan <- mulArgs{a: <-feedChan, b: <-feedChan}
		if log {
			fmt.Println("mulChan <- two numbers")
		}
	}

	defer close(mulChan)
	defer fmt.Println("\nDone!")
	return <-feedChan
}

func mulWorker(mulChan <-chan mulArgs, feedChan chan<- *big.Int) {
	for mulBig := range mulChan {
		feedChan <- mul(mulBig.a, mulBig.b)
		if log {
			fmt.Println("feedChan <- a×b")
		}
	}
}

func factorial(n uint64) *big.Int {
	fmt.Println("Digesting...")
	primePowers := factorize(n)
	powsLen := len(primePowers)

	numWorkers := runtime.GOMAXPROCS(0)
	powChan := make(chan powArgs)
	feedChan := make(chan *big.Int)
	mulChan := make(chan mulArgs, powsLen)

	fmt.Println("Using", numWorkers, "workers for powering and multiplication")
	for i := 0; i < numWorkers; i++ {
		go powWorker(powChan, feedChan)
		go mulWorker(mulChan, feedChan)
	}

	go initPowChan(primePowers, powChan)
	defer close(feedChan)

	return feedChanToMulChan(feedChan, mulChan, powsLen-1)
}
