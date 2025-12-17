package main

func getFirstPrimeFactor(x uint64, primes []uint64) uint64 {
	xSqrt := sqrt(x)
	for _, prime := range primes {
		if prime > xSqrt {
			break
		}
		if x%prime == 0 {
			return prime
		}
	}
	return x
}

func primeFactors(n uint64, primes []uint64) (res []uint64) {
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
