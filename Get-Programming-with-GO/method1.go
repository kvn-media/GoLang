package main

import "fmt"

type person struct {
	age int
}
// Nil receivers
/* 
func (p *person) birthday() {
	p.age++ // nil pointer dereference
} 
*/
// Guard clause
/* Go will happily call methods even when the receiver has a nil value. A nil receiver behaves no differently than a nil parameter. This means methods can guard against nil values, as shown in the following listing. Rather than check for nil before every call to the birthday method, the preceding listing guards against nil receivers inside the method.*/
func (p *person) birthday() {
	if p == nil {
	return
}
	p.age++
}

func main() {
	var nobody *person
	fmt.Println(nobody) // Prints <nil>
	nobody.birthday()

	// Function types that are nil
	var fn func(a, b int) int
	fmt.Println(fn == nil) // Prints True
}