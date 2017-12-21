package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"sync"
)

var big0 = big.NewInt(0)
var big1 = big.NewInt(1)

func factorial(x int) *big.Int {
	vals := make(chan *big.Int, x+1)
	res := make(chan *big.Int)
	count := x
	var countMutex sync.Mutex
	for i := 1; i <= x; i++ {
		j := big.NewInt(int64(i))
		vals <- j
		go func() {
			countMutex.Lock()
			count--
			countMutex.Unlock()
			if count == 0 {
				res <- <-vals
			} else {
				a := <-vals
				b := <-vals
				c := new(big.Int)
				c.Mul(a, b)
				vals <- c
			}
		}()
	}
	return <-res
}

func main() {
	if len(os.Args) != 2 {
		panic("We need exactly one argument which is numerical")
	}
	xStr := os.Args[1]
	x, parsingError := strconv.Atoi(xStr)
	if parsingError != nil {
		panic(fmt.Sprintf("Failed to parse %s %T as an integer number because %s", xStr, xStr, parsingError))
	}
	fmt.Printf("%d! = %d", x, factorial(x))
}
