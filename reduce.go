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
		if x, ok := red[ppi.power]; ok {
			red[ppi.power] = x.Mul(x, toBig(ppi.prime))
		} else {
			red[ppi.power] = toBig(ppi.prime)
		}
	}

	// fmt.Println("reduced", red)

	return red
}

func merge(pows PPBigTable, primes []uint64) PPBigTable {
	fmt.Println("Merging")
	m := make(PPBigTable)

	for power, prime := range pows {
		factors := primeFactors(power, primes)

		for _, factor := range factors {
			if x, ok := m[factor]; ok {
				// fmt.Println("Yes", power, prime, factor, x)
				var res big.Int
				m[factor] = res.Mul(x, prime)
			} else {
				// fmt.Println("no", power, prime, factor)
				m[factor] = prime
			}
		}
	}

	fmt.Println("Merged", len(pows), "to", len(m))

	return m
}
