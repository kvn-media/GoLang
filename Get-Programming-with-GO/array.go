package main

import "fmt"

func main() {
	var planets [8]string
	// Assigns a planet at index 0
	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2] // Retrieves the planet at index 2
	fmt.Println(earth) // Prints earth
	// The length of an array can be determined with the built-in len function. The other elements contain the zero value for their type, an empty string:
	fmt.Println(len(planets)) // Prints 8
	fmt.Println(planets[3] == "") // Prints true

	/* Invalid array index 8 [out of bounds for 8-element, array]
	var planets [8]string
	planets[8] = "Pluto"
	pluto := planets[8]
	// Panic: runtime error: index out of range
	var planets [8]string
	i := 8
	planets[i] = "Pluto"
	pluto := planets[i]
	*/
}