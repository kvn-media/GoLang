package main

import "fmt"

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	terrestrial := planets[0:4:4] // Length 4, capacity 4
	worlds := append(terrestrial, "Ceres")

	fmt.Println(planets, worlds) // Prints [Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]
	// terrestrial := planets[0:4] // Length 4, capacity 8
	fmt.Println(planets) // Prints [Mercury Venus Earth Mars Ceres Saturn Uranus Neptune]
}