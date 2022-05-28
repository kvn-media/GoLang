package main

import "fmt"

func main() {
	answer := 42
	address := &answer
	fmt.Printf("address is a %T\n", address) // Prints address is a *int

	/* The asterisk in *int denotes that the type is a pointer. In this case, it can point to other variables of type int.
	Pointer types can appear anywhere types are used, including in variable declarations, function parameters, return types, structure field types, and so on. 
	In the following listing, the asterisk (*) in the declaration of home indicates that it’s a pointer type. */
}