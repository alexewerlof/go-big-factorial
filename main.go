package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("We need exactly one argument")
		os.Exit(1)
	}
	xStr := os.Args[1]
	x, parsingError := strconv.Atoi(xStr)
	if parsingError != nil {
		fmt.Printf("Failed to parse %s %T as an integer number because %s\n", xStr, xStr, parsingError)
		os.Exit(1)
	}
	if x < 2 {
		fmt.Println("The number should be bigger than 1")
		os.Exit(1)
	}
	result := factorial(uint64(x))
	fmt.Println("Converting to string...")
	fmt.Printf("%d! = %d\n", x, result)
}
