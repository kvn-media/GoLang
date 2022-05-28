package main

import "fmt"

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
	planetsMarkII := planets // copies planet array
	planets[2] = "whoops" // Makes way for an interstellar bypass
	fmt.Println(planets) // Prints [Mercury Venus whoops Mars Jupiter Saturn Uranus Neptune]
	fmt.Println(planetsMarkII) // Prints [Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]
}