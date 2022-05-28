package main

import "fmt"

// func terraform(planets [8]string) {
// 	for i := range planets {
// 		planets[i] = "New" + planets[i]
// 	}
// }
// Planets attaches methods to []string.
type Planets []string

func (planets Planets) terraform() {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
}

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	Planets(planets[3:4]).terraform()
	Planets(planets[6:]).terraform()
	fmt.Println(planets)
	// terraform(planets)
	// fmt.Println(planets) // Prints [Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]
}