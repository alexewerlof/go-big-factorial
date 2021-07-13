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
	for6 := primeFactors(6, primes)
	if !sliceEq(for6, []uint64{2, 3}) {
		t.Error("Could not break 6 to its primes", for6)
	}

	for24 := primeFactors(24, primes)
	if !sliceEq(for24, []uint64{2, 2, 2, 3}) {
		t.Error("Could not break 24 to its primes", for24)
	}

	for13 := primeFactors(13, primes)
	if !sliceEq(for13, []uint64{13}) {
		t.Error("Could not break 13 to its primes", for13)
	}

	for2 := primeFactors(2, primes)
	if !sliceEq(for2, []uint64{2}) {
		t.Error("Could not break 2 to its primes", for2)
	}
}
