package main

import (
	"math/big"
	"testing"
)

func TestPpMapToArr(t *testing.T) {
	input := PPMap{
		31: 3,
		59: 1,
		83: 1,
		11: 9,
		17: 5,
		23: 4,
		79: 1,
		97: 1,
		13: 7,
		41: 2,
		43: 2,
		7:  16,
		37: 2,
		47: 2,
		53: 1,
		61: 1,
		2:  97,
		3:  48,
		5:  24,
		67: 1,
		73: 1,
		89: 1,
		19: 5,
		29: 3,
		71: 1,
	}

	expected := []PP{
		{53, 1},
		{59, 1},
		{61, 1},
		{67, 1},
		{71, 1},
		{73, 1},
		{79, 1},
		{83, 1},
		{89, 1},
		{97, 1},
		{37, 2},
		{41, 2},
		{43, 2},
		{47, 2},
		{29, 3},
		{31, 3},
		{23, 4},
		{17, 5},
		{19, 5},
		{13, 7},
		{11, 9},
		{7, 16},
		{5, 24},
		{3, 48},
		{2, 97},
	}

	result := ppMapToArr(input)

	if len(expected) != len(result) {
		t.Errorf("the lengths don't match. Expected %d but got %d", len(expected), len(result))
	}

	for i, r := range result {
		if expected[i] != r {
			t.Errorf("They are not equal. Expected %d but got %d", expected[i], r)
		}
	}
}

func (x *PPBigMap) deepEquals(y *PPBigMap, t *testing.T) {
	t.Helper()

	var lenX = len(*x)
	var lenY = len(*y)

	if lenX != lenY {
		t.Errorf("the lengths don't match. Expected %d but got %d", lenX, lenY)
	}

	for power, prime := range *y {
		if (*x)[power].Cmp(prime) != 0 {
			t.Errorf("Missing entry for power %d prime %d", power, prime)
		}
	}
}

func TestReduce(t *testing.T) {
	input := PPMap{
		31: 3,
		59: 1,
		83: 1,
		11: 9,
		17: 5,
		23: 4,
		79: 1,
		97: 1,
		13: 7,
		41: 2,
		43: 2,
		7:  16,
		37: 2,
		47: 2,
		53: 1,
		61: 1,
		2:  97,
		3:  48,
		5:  24,
		67: 1,
		73: 1,
		89: 1,
		19: 5,
		29: 3,
		71: 1,
	}

	expected := PPBigMap{
		1:  big.NewInt(53 * 59 * 61 * 67 * 71 * 73 * 79 * 83 * 89 * 97),
		2:  big.NewInt(37 * 41 * 43 * 47),
		3:  big.NewInt(29 * 31),
		4:  big.NewInt(23),
		5:  big.NewInt(17 * 19),
		7:  big.NewInt(13),
		9:  big.NewInt(11),
		16: big.NewInt(7),
		24: big.NewInt(5),
		48: big.NewInt(3),
		97: big.NewInt(2),
	}

	result := reduce(input)

	result.deepEquals(&expected, t)
}

func TestMerge(t *testing.T) {
	input := PPBigMap{
		1:  big.NewInt(53 * 59 * 61 * 67 * 71 * 73 * 79 * 83 * 89 * 97),
		2:  big.NewInt(37 * 41 * 43 * 47),
		3:  big.NewInt(29 * 31),
		4:  big.NewInt(23),
		5:  big.NewInt(17 * 19),
		7:  big.NewInt(13),
		9:  big.NewInt(11),
		16: big.NewInt(7),
		24: big.NewInt(5),
		48: big.NewInt(3),
		97: big.NewInt(2),
	}

	expected := PPBigMap{
		1:  big.NewInt(53 * 59 * 61 * 67 * 71 * 73 * 79 * 83 * 89 * 97),
		2:  big.NewInt(37 * 41 * 43 * 47 * 23 * 23 * 7 * 7 * 7 * 7 * 5 * 5 * 5 * 3 * 3 * 3 * 3),
		3:  big.NewInt(29 * 31 * 11 * 11 * 5 * 3),
		5:  big.NewInt(17 * 19),
		7:  big.NewInt(13),
		97: big.NewInt(2),
	}

	result := merge(input, primes)

	result.deepEquals(&expected, t)
}
