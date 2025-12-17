package main

import (
	"testing"
)

// deepEqualPPMap tests whether two PPMap structs contain the same elements.
func deepEqualPPMap(ppMap1, ppMap2 PriPow) bool {
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

func TestFactorize(t *testing.T) {
	knownResults := map[uint64]PriPow{
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
		result := factorize(n)
		if !deepEqualPPMap(result, expected) {
			t.Errorf("Could not factorize %d correctly\nExpected: %v\nGot:     %v", n, expected, result)
		}
	}
}
