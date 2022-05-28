package main

import "fmt"

func main() {

	type person struct {
		name, superpower string
		age              int
	}
	timmy := &person{
		name: "Timoty",
		age:  10,
	}

	superpowers := &[3]string{"flight", "invisibility", "super strength"}
	fmt.Println(superpowers[0])
	fmt.Println(superpowers[1:2])

	(*timmy).superpower = "flying"
	fmt.Printf("%+v\n", timmy)
}