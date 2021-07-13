package main

import (
	"fmt"
	"math/big"
	"sort"
)

type PPBigMap map[uint64]*big.Int

func ppMapToArr(ppMap PPMap) []PP {
	pp := make([]PP, 0, len(ppMap))

	for prime, power := range ppMap {
		// fmt.Println(prime, power)
		pp = append(pp, PP{prime, power})
	}
	// fmt.Println("pp", pp)

	sort.Slice(pp, func(i, j int) bool {
		if pp[i].power == pp[j].power {
			return pp[i].prime < pp[j].prime
		}
		return pp[i].power < pp[j].power
	})

	// fmt.Println("pp", pp)
	return pp
}

func (p *PPBigMap) update(power uint64, prime *big.Int) bool {
	if x, ok := (*p)[power]; ok {
		var res big.Int
		(*p)[power] = res.Mul(x, prime)
		return true
	}

	(*p)[power] = prime

	return false
}

func reduce(pows PPMap) PPBigMap {
	fmt.Println("Reducing")

	pp := ppMapToArr(pows)
	red := make(PPBigMap)

	for _, ppi := range pp {
		red.update(ppi.power, toBig(ppi.prime))
	}

	// fmt.Println("reduced", red)

	return red
}

func merge(ppBigMap PPBigMap, primes []uint64) PPBigMap {
	fmt.Println("Merging")
	m := make(PPBigMap)

	for power, prime := range ppBigMap {
		factors := primeFactors(power, primes)

		for _, factor := range factors {
			m.update(factor, prime)
		}
	}

	fmt.Println("Merged", len(ppBigMap), "to", len(m), m)

	return m
}

func allMergedReducedPrimeFactors(n uint64) PPBigMap {
	pows, primes := factorize(n)

	return merge(reduce(pows), primes)
}
