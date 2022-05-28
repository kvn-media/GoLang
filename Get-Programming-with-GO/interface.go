package main

import "fmt"

func main() {
	var v interface{}
	fmt.Printf("%T %v %v\n", v, v, v == nil) // Prints <nil> <nil> true
	/* When a variable with an interface type is assigned a value, the interface internally points to the type and value of that variable. This leads to the rather surprising behavior of a nil value that doesn’t compare as equal to nil. Both the interface type and value need to be nil for the variable to equal nil, as shown in the following listing. */
	var p *int
	v = p
	fmt.Printf("%T %v %v\n", v, v, v == nil) // Prints *int <nil> false
	/* The %#v format verb is shorthand to see both type and value, also revealing that the variable contains (*int)(nil) rather than just<nil> */
	fmt.Printf("%#v\n", v) // Prints (*int)(nil)
}