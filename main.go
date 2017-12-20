package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var big0 = big.NewInt(0)
var big1 = big.NewInt(1)

func factorial(x int) *big.Int {
	vals := make(chan *big.Int, x+1)
	res := make(chan *big.Int)
	count := make(chan int)
	for i := 1; i <= x; i++ {
		j := big.NewInt(int64(i))
		vals <- j
		go func() {
			c := <-count
			if c == 0 {
				res <- <-vals
			} else {
				a := <-vals
				b := <-vals
				count <- c - 1
				c := new(big.Int)
				c.Mul(a, b)
				vals <- c
			}
		}()
	}
	count <- x - 1
	return <-res
}

func main() {
	if len(os.Args) != 2 {
		panic("We need exactly one argument which is numerical")
	}
	nStr := os.Args[1]
	nParsed, parsingError := strconv.Atoi(nStr)
	if parsingError != nil {
		panic(fmt.Sprintf("Failed to parse %s %T as an integer number because %s", nStr, nStr, parsingError))
	}
	fmt.Printf("%d!\n", nParsed)
	fmt.Println(factorial(nParsed))
}
