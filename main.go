package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"sync"
)

type opCounter struct {
	val   int
	mutex sync.Mutex
}

func (cnt *opCounter) dec() (isNotZero bool) {
	cnt.mutex.Lock()
	cnt.val--
	defer cnt.mutex.Unlock()
	isNotZero = cnt.val != 0
	return
}

func factorial(x int) *big.Int {
	vals := make(chan *big.Int, x+1)
	res := make(chan *big.Int)
	count := opCounter{val: x}
	for i := 1; i <= x; i++ {
		val := big.NewInt(int64(i))
		vals <- val
	}
	go func() {
		for count.dec() {
			c := new(big.Int)
			c.Mul(<-vals, <-vals)
			vals <- c
		}
		res <- <-vals
	}()
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
	fmt.Printf("%d! = %d\n", x, factorial(x))
}
