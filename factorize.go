package main

import (
	"math"
)

func getFirstPrimeFactor(x uint64, primes []uint64) uint64 {
	xSqrt := uint64(math.Sqrt(float64(x)))
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

// PriPow is a map of prime factors (keys) and their power (value)
type PriPow map[uint64]uint64

func factorize(n uint64) PriPow {
	var primes []uint64
	var m = make(PriPow)
	for i := uint64(2); i <= n; i++ {
		pFactors := primeFactors(i, primes)
		if len(pFactors) == 1 {
			// Oh we found a new prime!
			primes = append(primes, pFactors[0])
		}
		for _, pFactor := range pFactors {
			m[pFactor]++
		}
	}

	return m
}
