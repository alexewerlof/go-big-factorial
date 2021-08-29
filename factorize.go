package main

import (
	"runtime"
)

func getFirstPrimeFactor(x uint64, primes []uint64) uint64 {
	xSqrt := sqrt(x)
	for _, prime := range primes {
		if prime > xSqrt {
			break
		}
		if x%prime == 0 {
			return prime
		}
	}
	return x
}

func primeFactors(n uint64, primes []uint64) (res []uint64) {
	x := n
	for {
		p := getFirstPrimeFactor(x, primes)
		res = append(res, p)
		if x == p {
			break
		}
		x /= p
	}
	return
}

func factorizeWorker(primes []uint64, nCh <-chan uint64, factorsCh chan<- []uint64) {
	for n := range nCh {
		factorsCh <- primeFactors(n, primes)
	}
}

func feedNCh(n uint64, nCh chan<- uint64) {
	for i := uint64(2); i <= n; i++ {
		nCh <- i
	}
	close(nCh)
}

// PriPow is a map of prime factors (keys) and their power (value)
type PriPow map[uint64]uint64

func factorize(n uint64) PriPow {
	primes := primesUnder(sqrt(n))
	numWorkers := runtime.GOMAXPROCS(0)

	nCh := make(chan uint64)
	factorsCh := make(chan []uint64)
	defer close(factorsCh)

	go feedNCh(n, nCh)

	for i := 0; i < numWorkers; i++ {
		go factorizeWorker(primes, nCh, factorsCh)
	}

	m := make(PriPow)

	// We feed nCh from 2, therefore the number of results is one less than n as well
	for i := uint64(1); i < n; i++ {
		factors := <-factorsCh
		for _, f := range factors {
			m[f]++
		}
	}

	return m
}
