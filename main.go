package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

const log = true

func toBig(x uint64) *big.Int {
	return big.NewInt(int64(x))
}

func pow(a, b uint64, vals chan<- *big.Int) {
	var res big.Int
	if log {
		fmt.Println(a, "^", b)
	}
	vals <- res.Exp(toBig(a), toBig(b), nil)
}

func mul(vals chan *big.Int, done chan<- bool) {
	var res big.Int
	if log {
		fmt.Println("x")
	}
	vals <- res.Mul(<-vals, <-vals)
	done <- true
}

func factorial(n uint64) *big.Int {
	fmt.Println("Digesting...")
	primePowers := digestAllUnder(n)
	powsLen := len(primePowers)
	vals := make(chan *big.Int, powsLen)
	//fmt.Println("Digested to", pows)
	fmt.Println("\nPowering...")
	for prime, power := range primePowers {
		//fmt.Println("a=", a, "b=", b, "pow=", pow(a, b))
		if power == 1 {
			vals <- toBig(prime)
		} else {
			go pow(prime, power, vals)
		}
	}
	primePowers = nil
	fmt.Println("\nMultiplying...")
	multiOp := make(chan bool)
	for i := 1; i < powsLen; i++ {
		go mul(vals, multiOp)
	}
	for i := 1; i < powsLen; i++ {
		x := <-multiOp
		if log {
			fmt.Println("One multiplication operation was done", i, x)
		}
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
	fmt.Printf("%d! = %d\n", x, factorial(uint64(x)))
}
