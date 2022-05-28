package main

import "fmt"

var f = func()  { // Assigns an anonymous function to a variable
	fmt.Println("Dress up for a masquerade.")
}

func main() {
	f() // Prints Dress up for the masquerade.
}