package main

import (
	"fmt"
	"math/big"
	"sort"
)

type PPBigTable map[uint64]*big.Int

func powTableToPPArr(pows PPTable) []PP {
	pp := make([]PP, 0, len(pows))
	i := 0
	for prime, power := range pows {
		// fmt.Println(prime, power)
		pp = append(pp, PP{prime, power})
		i++
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

func reduce(pows PPTable) PPBigTable {
	fmt.Println("Reducing")
	pp := powTableToPPArr(pows)
	red := make(PPBigTable)
	for _, ppi := range pp {
		red.setEl(ppi.power, toBig(ppi.prime))
	}

	// fmt.Println("reduced", red)

	return red
}

func (p *PPBigTable) setEl(power uint64, prime *big.Int) bool {
	if x, ok := (*p)[power]; ok {
		var res big.Int
		(*p)[power] = res.Mul(x, prime)
		return true
	}
	(*p)[power] = prime
	return false
}

func merge(pows PPBigTable, primes []uint64) PPBigTable {
	fmt.Println("Merging")
	m := make(PPBigTable)

	for power, prime := range pows {
		factors := primeFactors(power, primes)

		for _, factor := range factors {
			m.setEl(factor, prime)
		}
	}

	fmt.Println("Merged", len(pows), "to", len(m), m)

	return m
}

func allMergedReducedPrimeFactors(n uint64) PPBigTable {
	pows, primes := allPrimeFactors(n)

	return merge(reduce(pows), primes)
}
