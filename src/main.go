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

func pow(a, b int64) (res *big.Int) {
	res = new(big.Int)
	res.Exp(big.NewInt(a), big.NewInt(b), nil)
	return
}

func mul(a, b *big.Int) (res *big.Int) {
	res = new(big.Int)
	res.Mul(a, b)
	return
}

func factorial(x int64) *big.Int {
	fmt.Println("Digesting...")
	pows := digestAll(x)
	vals := make(chan *big.Int, len(pows))
	//fmt.Println("Digested to", pows)
	fmt.Println("Powering...")
	for a, b := range pows {
		//fmt.Println("a=", a, "b=", b, "pow=", pow(a, b))
		fmt.Print(".")
		vals <- pow(a, b)
	}
	pows = nil

	/*
		for i := 1; i < len(pows); i++ {
			vals <- mul(<-vals, <-vals)
		}
		return <-vals
	*/
	fmt.Println("Multiplying")
	var workers workerCounter
	for c := len(vals); c > 1; c = workers.count + len(vals) {
		workers.mutex.Lock()
		workers.count++
		workers.mutex.Unlock()
		go func(a, b *big.Int) {
			res := mul(a, b)
			workers.mutex.Lock()
			vals <- res
			workers.count--
			workers.mutex.Unlock()
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
