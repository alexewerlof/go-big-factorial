package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		panic("We need exactly one argument")
	}
	xStr := os.Args[1]
	x, parsingError := strconv.Atoi(xStr)
	if parsingError != nil {
		panic(fmt.Sprintf("Failed to parse %s %T as an integer number because %s", xStr, xStr, parsingError))
	}
	if x < 2 {
		panic("The number should be bigger than 1")
	}
	result := factorial(uint64(x))
	fmt.Println("Converting to string...")
	fmt.Printf("%d! = %d\n", x, result)
}
