package main

import (
	"fmt"
	"math/big"
	"os"
)

var big0 = big.NewInt(0)
var big1 = big.NewInt(1)

func factorial(x *big.Int) *big.Int {
	if x.Cmp(big0) == 0 {
		return big1
	}
	var res, xMinOne big.Int
	return res.Mul(x, factorial(xMinOne.Sub(x, big1)))
}

func main() {
	if len(os.Args) != 2 {
		panic("We need exactly one argument which is numerical")
	}
	nStr := os.Args[1]
	nParsed := new(big.Int)
	_, parsingSuccess := nParsed.SetString(nStr, 10)
	if !parsingSuccess {
		panic(fmt.Sprintf("Failed to parse %s %T as an integer number", nStr, nStr))
	}
	fmt.Printf("%d!\n", nParsed)
	fmt.Println(factorial(nParsed))
}
