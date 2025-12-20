package main

import (
	"fmt"
	"strconv"

	"rsc.io/quote"
)

func Hello() string {
	return "Hello - World!"
}

func main() {
	// Example of printing "Hello World"
	fmt.Println("Hello World")
	// Example of using the quote package to print a quote
	fmt.Println(quote.Go())
	//
	// Example of using strconv to convert an integer to a string
	num := 42
	fmt.Printf("%v, %T\n", num, num)
	// Convert integer to string
	strNum := strconv.Itoa(num)
	fmt.Printf("%v, %T", strNum, strNum)
}
