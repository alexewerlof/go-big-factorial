package main

import (
	"fmt"
	"math"
)

func getFirstPrimeFactor(x int64, primes *[]int64) int64 {
	limit := math.Sqrt(float64(x))
	//fmt.Println("x=", x)
	for _, p := range *primes {
		//fmt.Println("  p=", p)
		if float64(p) > limit {
			return x
		}
		if math.Mod(float64(x), float64(p)) == 0 {
			return p
		}
	}
	return x
}

func digest(n int64, primes *[]int64) (res []int64) {
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

func digestAll(n int64) (pows map[int64]int64) {
	var primes []int64
	pows = make(map[int64]int64)
	for i := int64(2); i <= n; i++ {
		res := digest(i, &primes)
		fmt.Print(":")
		if len(res) == 1 {
			primes = append(primes, res[0])
			//fmt.Println("Oh a new prime:", res, primes)
		}
		for _, r := range res {
			rInt := int64(r)
			v, exists := pows[rInt]
			if exists {
				pows[rInt] = v + 1
			} else {
				pows[rInt] = 1
			}
		}
		//fmt.Println("digest", i, res)
	}
	return pows
}
