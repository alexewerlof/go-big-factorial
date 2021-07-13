package main

import (
	"testing"
)

// Prime numbers under 100
var primes = []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

// sliceEq tells whether a and b contain the same elements.
func sliceEq(slice1, slice2 []uint64) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}

	return true
}

func TestPrimeFactors(t *testing.T) {
	knownResults := map[uint64][]uint64{
		2:  {2},
		6:  {2, 3},
		13: {13},
		24: {2, 2, 2, 3},
	}

	for n, expected := range knownResults {
		result := primeFactors(6, primes)
		if !sliceEq(result, expected) {
			t.Errorf("Could not break %d to its primes\nExpected: %v\nGot:     %v", n, expected, result)
		}
	}
}
