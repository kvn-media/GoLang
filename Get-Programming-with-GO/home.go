package main

import "fmt"

func main() {
	canada := "Canada"

	var home *string
	fmt.Printf("home is a %T\n", home) // Prints home is a *string

	home = &canada
	fmt.Println(*home) // Print Canada
	/* TIP An asterisk prefixing a type denotes a pointer type, whereas an asterisk prefixing a variable name is used to dereference the value that variable points to.
	
	The home variable in the previous listing can point at any variable of type string. However, the Go compiler won’t allow home to point to a variable of any other type, such as int. */
}