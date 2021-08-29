package main

import (
	"math"
)

func isPrime(x uint64, primes []uint64) bool {
	xSqrt := uint64(math.Sqrt(float64(x)))

	for _, prime := range primes {
		if prime > xSqrt {
			break
		}
		if x%prime == 0 {
			return false
		}
	}
	return true
}

func primesUnder(max uint64) []uint64 {
	primes := []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	for i := uint64(101); i < max; i += 2 {
		if i%5 == 0 {
			continue
		}
		if isPrime(i, primes) {
			primes = append(primes, i)
		}
	}

	return primes
}
