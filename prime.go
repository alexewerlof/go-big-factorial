package main

import (
	"math"
)

func isDivisible(x, prime float64) bool {
	return math.Mod(x, prime) == 0
}

func getFirstPrimeFactor(x uint64, primes []uint64) uint64 {
	xF64 := float64(x)
	limit := math.Sqrt(xF64)
	//fmt.Println("x=", x)
	for _, p := range primes {
		pF64 := float64(p)
		//fmt.Println("  p=", p)
		if pF64 > limit {
			break
		}
		if isDivisible(xF64, pF64) {
			return p
		}
	}
	return x
}

func digest(n uint64, primes []uint64) (res []uint64) {
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

func digestAllUnder(n uint64) (pows map[uint64]uint64) {
	var primes []uint64
	pows = make(map[uint64]uint64)
	for i := uint64(2); i <= n; i++ {
		res := digest(i, primes)
		//fmt.Print(":")
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
	return pows
}
