package main

import (
	"fmt"
	"math/big"
	"runtime"
)

const log = true

func initPowChan(primePowers PPMap, powChan chan<- PowArgs) {
	for x, n := range primePowers {
		powChan <- PowArgs{toBig(x), n}
	}
}

func powWorker(powChan <-chan PowArgs, feedChan chan<- *big.Int) {
	for ppBig := range powChan {
		feedChan <- pow(ppBig)
		if log {
			fmt.Println("feedChan <- x^y")
		}
	}
}

func feedChanToMulChan(feedChan <-chan *big.Int, mulChan chan<- MulArgs, times int) *big.Int {
	for i := 0; i < times; i++ {
		mulChan <- MulArgs{a: <-feedChan, b: <-feedChan}
		if log {
			fmt.Println("mulChan <- two numbers")
		}
	}

	defer fmt.Println("\nDone!")
	return <-feedChan
}

func mulWorker(mulChan <-chan MulArgs, feedChan chan<- *big.Int) {
	for mulBig := range mulChan {
		feedChan <- mul(mulBig)
		if log {
			fmt.Println("feedChan <- a×b")
		}
	}
}

func factorial(n uint64) *big.Int {
	fmt.Println("Digesting...")
	primePowers := factorize(n)
	powsLen := len(primePowers)
	//fmt.Println("Digested to", pows)

	numWorkers := runtime.GOMAXPROCS(0)
	powChan := make(chan PowArgs)
	feedChan := make(chan *big.Int)
	mulChan := make(chan MulArgs, powsLen)

	fmt.Println("Using", numWorkers, "workers for powering and multiplication")
	for i := 0; i < numWorkers; i++ {
		go powWorker(powChan, feedChan)
		go mulWorker(mulChan, feedChan)
	}

	go initPowChan(primePowers, powChan)

	return feedChanToMulChan(feedChan, mulChan, powsLen-1)
}
