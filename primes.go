package main

func isPrime(x uint64, primes []uint64) bool {
	xSqrt := sqrt(x)

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
	initialPrimes := []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	var primes []uint64
	for _, p := range initialPrimes {
		if p >= max {
			return primes
		}
		primes = append(primes, p)
	}

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
