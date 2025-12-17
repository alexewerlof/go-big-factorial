package main

// PriPow is a map of prime factors (keys) and their power (value)
type PriPow map[uint64]uint64

func factorize(n uint64) PriPow {
	// Legendre's formula requires primes up to n
	primes := primesUnder(n + 1)
	m := make(PriPow)

	for _, p := range primes {
		count := uint64(0)
		// Sum floor(n / p^k)
		for div := p; div <= n; div *= p {
			count += n / div
			// Avoid overflow for next iteration
			if n/p < div {
				break
			}
		}
		if count > 0 {
			m[p] = count
		}
	}

	return m
}
