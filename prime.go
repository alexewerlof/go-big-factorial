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

type PPTable map[uint64]uint64

func allPrimeFactors(n uint64) (PPTable, []uint64) {
	var primes []uint64
	var pows = make(PPTable)
	for i := uint64(2); i <= n; i++ {
		res := primeFactors(i, primes)
		//fmt.Print("/")
		if len(res) == 1 {
			primes = append(primes, res[0])
			//fmt.Println("Oh a new prime:", res, primes)
		}
		for _, r := range res {
			rInt := uint64(r)
			_, exists := pows[rInt]
			if exists {
				pows[rInt]++
			} else {
				pows[rInt] = 1
			}
		}
		//fmt.Println("digest", i, res)
	}
	return pows, primes
}
