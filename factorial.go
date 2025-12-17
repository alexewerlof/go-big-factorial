package main

import (
	"fmt"
	"math/big"
	"runtime"
	"sync"
)

const log = false

func productTree(nums []*big.Int) *big.Int {
	if len(nums) == 0 {
		return big.NewInt(1)
	}
	if len(nums) == 1 {
		return nums[0]
	}
	mid := len(nums) / 2
	return mul(productTree(nums[:mid]), productTree(nums[mid:]))
}


func factorial(n uint64) *big.Int {
	fmt.Println("Digesting...")
	primePowers := factorize(n)

	numWorkers := runtime.GOMAXPROCS(0)
	fmt.Println("Using", numWorkers, "workers for powering")

	// Input channel for jobs
	type job struct {
		p, k uint64
	}
	jobs := make(chan job, len(primePowers))
	results := make(chan *big.Int, len(primePowers))

	// Spawn workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				results <- pow(toBig(j.p), j.k)
			}
		}()
	}

	// Send jobs
	for p, k := range primePowers {
		jobs <- job{p, k}
	}
	close(jobs)

	// Wait for workers
	wg.Wait()
	close(results)

	// Collect results
	nums := make([]*big.Int, 0, len(primePowers))
	for res := range results {
		nums = append(nums, res)
	}

	fmt.Println("Multiplying...")
	return productTree(nums)
}
