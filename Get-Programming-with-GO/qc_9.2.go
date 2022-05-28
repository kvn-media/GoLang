package main

import "fmt"

func main() {
	var star byte = '*'
	fmt.Printf("%c %[1]v\n", star) // Prints * 42

	smile := ''
	fmt.Printf("%c %[1]v\n", smile) // Prints  128515

	acute :=  'é'
	fmt.Printf("%c %[1]v\n", acute) // Prints é 233
}