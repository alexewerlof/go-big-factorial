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
	powsLen := len(pows)
	vals := make(chan *big.Int, powsLen)
	//fmt.Println("Digested to", pows)
	fmt.Println("\nPowering...")
	for a, b := range pows {
		//fmt.Println("a=", a, "b=", b, "pow=", pow(a, b))
		fmt.Print("^")
		vals <- pow(a, b)
	}
	pows = nil
	fmt.Println("\nMultiplying...")
	for i := 1; i < powsLen; i++ {
		fmt.Print("x")
		vals <- mul(<-vals, <-vals)
	}
	fmt.Println("\nDone!")
	defer fmt.Println("Converting to string...")
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
