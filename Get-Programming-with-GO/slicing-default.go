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
	terrestial := planets[:4]
	gasGiats := planets[4:6]
	iceGiants := planets[6:]
	fmt.Println(terrestial, gasGiats, iceGiants)
}
