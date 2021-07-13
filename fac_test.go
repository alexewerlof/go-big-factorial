package main

import (
	"math/big"
	"testing"
)

func TestFactorial5(t *testing.T) {
	expected := big.NewInt(120)
	result := factorial(5)

	if result.Cmp(expected) != 0 {
		t.Error("Failed to compute 5!", result, expected)
	}
}

func TestFactorial99(t *testing.T) {
	result := factorial(100)

	if result.String() != "933262154439441526816992388562667004907159682643816214685929638952175999932299156089414639761565182862536979208272237582511852109168640000000000000000000000" {
		t.Error("Failed to compute 99!")
	}
}
