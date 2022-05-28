package main

import (
	"fmt"
)

func main() {
	third := 1.0 / 3
	fmt.Println(third) // Prints 0.3333333333333333
	fmt.Printf("%v\n", third) // Prints 0.3333333333333333
	fmt.Printf("%v\n", third) // Prints 0.333333
	fmt.Printf("%.3f\n", third) // Prints 0.333
	fmt.Printf("%4.2f\n", third) // Prints 0.33

	// The %f verb formats the value of third with a width and with precision
	// Precision specifies how many digits should appear after the decimal point; two digits for %.2f, for example

	fmt.Printf("%05.2f\n", third) // Prints 00.33

	/*
	third := 1.0 / 3
	fmt.Printf("%f\n", third) // Prints 0.333333
	fmt.Printf("%7.4f\n", third) // Prints 0.3333
	fmt.Printf("%06.2f\n", third) // Prints 000.33
	*/
}