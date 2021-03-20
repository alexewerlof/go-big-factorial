package main

import (
	"math/big"
	"testing"
)

func TestFactorial(t *testing.T) {
	expected := big.NewInt(120)
	result := factorial(5)

	if result.Cmp(expected) != 0 {
		t.Error("Failed to compute 5!", result, expected)
	}
}
