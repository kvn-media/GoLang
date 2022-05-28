package main

import "fmt"

func main() {

	piggyBank := 0.0
		for i := 0; i < 11; i++ {
 		piggyBank += 0.1
	}
	fmt.Println(piggyBank) // Prints 1.0999999999999999
	/*
	celsius := 21.0
	fmt.Println((celsius/5.0*9.0)+32, "º F\n") // Prints 69.80000000000001º F
	fmt.Println((9.0/5.0*celsius)+32, "º F\n") // Prints 69.80000000000001º F
	*/
}