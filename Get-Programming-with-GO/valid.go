package main

import "fmt"

type number struct {
	value int
	valid bool
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}
func (n number) String() string {
	if !n.valid {
		return "not set"
	}
	return fmt.Sprintf("%d", n.value)
}
func main() {
	n := newNumber(42)
	fmt.Println(n) // Prints 42

	e := number{}
	fmt.Println(e) // Prints not set
}
/*
 Nil pointer dereferences will crash your program.
 Methods can guard against receiving nil values.
 Default behavior can be provided for functions passed as arguments.
 A nil slice is often interchangeable with an empty slice.
 A nil map can be read from but not written to.
 If an interface looks like it’s nil, be sure both the type and value are nil.
 Nil isn’t the only way to represent nothing.
*/