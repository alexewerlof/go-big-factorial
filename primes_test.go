package main

import (
	"reflect"
	"testing"
)

func TestPrimesUnder(t *testing.T) {
	res := primesUnder(110)
	expected := []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109}
	if !reflect.DeepEqual(res, expected) {
		t.Error("Got unexpected results.\n", res, "\n", expected)
	}
}
