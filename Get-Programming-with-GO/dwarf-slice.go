package main

import "fmt"

func main() {
	// If you need a slice that reveals every element of the underlying array, one option is to declare an array and then slice it with [:], like this:
	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfSlice := dwarfArray[:]
	// Declare a slice directly
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Printf("array %T\n", dwarfArray)
	fmt.Printf("slice %T\n", dwarfs)
	fmt.Printf("dwarf %T\n", dwarfSlice)
}