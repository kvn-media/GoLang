package main

import "fmt"

func main() {
	answer := 42
	fmt.Println(&answer) // Prints 0x1040c108

	address := &answer
	fmt.Println(*address) // 42
	/* It prints 42 because the memory address (&) is dereferenced (*) back to the value. 
	Multiplication is an infix operator requiring two values, whereas dereferencing prefixes a single variable.*/
}