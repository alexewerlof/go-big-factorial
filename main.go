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

func worker(count *opCounter, vals chan *big.Int, res chan<- *big.Int) {
	for count.dec() {
		c := new(big.Int)
		c.Mul(<-vals, <-vals)
		vals <- c
	}
	res <- <-vals
}

func factorial(x int) *big.Int {
	vals := make(chan *big.Int, x+1)
	res := make(chan *big.Int)
	count := opCounter{val: x - 1}
	for i := 2; i <= x; i++ {
		val := big.NewInt(int64(i))
		vals <- val
	}
	go worker(&count, vals, res)
	return <-res
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
	fmt.Printf("%d! = %d\n", x, factorial(x))
}
