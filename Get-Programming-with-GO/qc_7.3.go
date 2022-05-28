package main

import "fmt"

func main() {
	a := "text"
	fmt.Printf("Type %T for %[1]v\n", a) // Prints Type string for text
	b := 42
	fmt.Printf("Type %T for %[1]v\n", b) // Prints Type int for 42
	c := 3.14
	fmt.Printf("Type %T for %[1]v\n", c) // Prints Type float64 for 3.14
	d := true
	fmt.Printf("Type %T for %[1]v\n", d) // Prints Type bool for true
}