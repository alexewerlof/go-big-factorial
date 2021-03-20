package main

import (
	"fmt"
	"math"
	"math/big"
	"sort"
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

type PP struct {
	prime uint64
	power uint64
}

func powTableToPPArr(pows map[uint64]uint64) []PP {
	pp := make([]PP, 0, len(pows))
	i := 0
	for prime, power := range pows {
		fmt.Println(prime, power)
		pp = append(pp, PP{prime, power})
		i++
	}
	fmt.Println("pp", pp)

	sort.Slice(pp, func(i, j int) bool {
		if pp[i].power == pp[j].power {
			return pp[i].prime < pp[j].prime
		}
		return pp[i].power < pp[j].power
	})

	fmt.Println("pp", pp)
	return pp
}

func reducer(pows map[uint64]uint64) map[uint64]*big.Int {
	pp := powTableToPPArr(pows)
	red := make(map[uint64]*big.Int)
	for _, ppi := range pp {
		if x, ok := red[ppi.power]; ok {
			red[ppi.power] = x.Mul(x, toBig(ppi.prime))
		} else {
			red[ppi.power] = toBig(ppi.prime)
		}
	}

	fmt.Println("red", red)

	return red
}

func digestAllUnder(n uint64) (pows map[uint64]uint64) {
	var primes []uint64
	pows = make(map[uint64]uint64)
	for i := uint64(2); i <= n; i++ {
		res := digest(i, primes)
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
	return pows
}
