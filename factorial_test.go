package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	knownFactorials := map[uint64]string{
		2: "2",
		5: "120",
		// Use: https://www.wolframalpha.com/
		10: "3628800",
		20: "2432902008176640000",
		99: "933262154439441526816992388562667004907159682643816214685929638952175999932299156089414639761565182862536979208272237582511852109168640000000000000000000000",
	}

	for n, expected := range knownFactorials {

		result := factorial(n)

		if result.String() != expected {
			t.Errorf("Failed to compute %d!\nExpected: %s\nGot:      %s", n, expected, result)
		}
	}
}
