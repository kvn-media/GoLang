package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

func main() {
	rebecca := person{
		name:       "Rebecca",
		superpower: "Imagination",
		age:        14,
	}
	birthday(&rebecca)
	fmt.Printf("%+v\n", rebecca) // Prints {name:Rebecca superpower:imagination age:15}
}