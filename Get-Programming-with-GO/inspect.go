package main

import "fmt"

func main() {
	year := 2018
	fmt.Printf("Type %T for %v\n", year, year) // Prints Type int for 2018

	days := 365.2425
	fmt.Printf("Type %T for %[1]v\n", days) // Prints Type float64 for 365.2425
}