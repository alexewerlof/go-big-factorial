package main

import (
	"testing"
)

// Prime numbers under 100
var primes = []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

// deepEqualSliceUint64 tests whether two slices contain the same elements.
func deepEqualSliceUint64(slice1, slice2 []uint64) bool {
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

// deepEqualPPMap tests whether two PPMap structs contain the same elements.
func deepEqualPPMap(ppMap1, ppMap2 PPMap) bool {
	if len(ppMap1) != len(ppMap2) {
		return false
	}

	for k, v := range ppMap1 {
		if v != ppMap2[k] {
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
		result := primeFactors(n, primes)
		if !deepEqualSliceUint64(result, expected) {
			t.Errorf("Could not break %d to its primes\nExpected: %v\nGot:     %v", n, expected, result)
		}
	}
}

func TestFactorize(t *testing.T) {
	knownResults := map[uint64]PPMap{
		2: {2: 1},
		10: {
			7: 1,
			5: 2,
			3: 4,
			2: 8,
		},
		13: {
			13: 1,
			11: 1,
			7:  1,
			5:  2,
			3:  5,
			2:  10,
		},
	}

	for n, expected := range knownResults {
		result, _ := factorize(n)
		if !deepEqualPPMap(result, expected) {
			t.Errorf("Could not factorize %d correctly\nExpected: %v\nGot:     %v", n, expected, result)
		}
	}
}
