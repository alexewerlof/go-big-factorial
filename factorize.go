package main

import (
	"math"
)

func getFirstPrimeFactor(x uint64, primes []uint64) uint64 {
	xSqrt := uint64(math.Sqrt(float64(x)))
	//fmt.Println("x=", x)
	for _, prime := range primes {
		//fmt.Println("  p=", p)
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
	// nSqr := math.Sqrt(n)
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

type PP struct {
	prime uint64
	power uint64
}

type PPMap map[uint64]uint64

func factorize(n uint64) (PPMap, []uint64) {
	var primes []uint64
	var pows = make(PPMap)
	for i := uint64(2); i <= n; i++ {
		pFactors := primeFactors(i, primes)
		//fmt.Print("/")
		if len(pFactors) == 1 {
			primes = append(primes, pFactors[0])
			//fmt.Println("Oh a new prime:", res, primes)
		}
		for _, r := range pFactors {
			pows[r]++
		}
		//fmt.Println("digest", i, res)
	}
	return pows, primes
}
