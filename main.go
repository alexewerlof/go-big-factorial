package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"sync"
)

type workerCounter struct {
	count int
	mutex sync.Mutex
}

func (workers *workerCounter) change(diff int) {
	workers.mutex.Lock()
	workers.count += diff
	workers.mutex.Unlock()
}

func factorial(x int64) *big.Int {
	vals := make(chan *big.Int, x)
	for i := int64(2); i <= x; i++ {
		vals <- big.NewInt(i)
	}
	var workers workerCounter
	for c := len(vals); c > 1; c = workers.count + len(vals) {
		workers.change(1)
		go func(a, b *big.Int) {
			res := new(big.Int)
			workers.change(-1)
			vals <- res.Mul(a, b)
		}(<-vals, <-vals)
	}
	return <-vals
}

func main() {
	if len(os.Args) != 2 {
		panic("We need exactly one argument which should be a number bigger than 2")
	}
	xStr := os.Args[1]
	x, parsingError := strconv.Atoi(xStr)
	if parsingError != nil {
		panic(fmt.Sprintf("Failed to parse %s %T as an integer number because %s", xStr, xStr, parsingError))
	}
	if x < 2 {
		panic("The number should be bigger than 2")
	}
	fmt.Printf("%d! = %d\n", x, factorial(int64(x)))
}
