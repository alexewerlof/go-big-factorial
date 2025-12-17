package main

func primesUnder(max uint64) []uint64 {
	if max <= 2 {
		return nil
	}
	// Sieve of Eratosthenes
	isPrime := make([]bool, max)
	for i := range isPrime {
		isPrime[i] = true
	}
	// 0 and 1 are not primes
	// We start loop from 2.
	for p := uint64(2); p*p < max; p++ {
		if isPrime[p] {
			for i := p * p; i < max; i += p {
				isPrime[i] = false
			}
		}
	}
	var primes []uint64
	for i := uint64(2); i < max; i++ {
		if isPrime[i] {
			primes = append(primes, uint64(i))
		}
	}
	return primes
}
