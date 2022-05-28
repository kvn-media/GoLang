package main

import "fmt"

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris", "Orcus"}

	// The append function is variadic, like Println, so you can pass multiple elements to append in one go:
	dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")
	// The slice contains nine dwarf planets, which can be determined with the len built-in function:
	fmt.Println(len(dwarfs)) // Prints [Ceres Pluto Haumea Makemake Eris Orcus Salacia Quaoar Sedna]
	// Prints 9
}