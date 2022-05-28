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
	terrestial := planets[0:4]
	gasGiats := planets[4:6]
	iceGiats := planets[6:8]
	fmt.Println(terrestial, gasGiats, iceGiats) // Prints [Mercury Venus Earth Mars] [Jupiter Saturn] [Uranus Neptune]
	fmt.Println(gasGiats[0]) // Prints Jupiter

	giants := planets[4:8]
	gas := giants[0:2]
	ice := giants[2:4]
	fmt.Println(giants, gas, ice) // Prints [Jupiter Saturn Uranus Neptune] [Jupiter Saturn] [Uranus Neptune]

	iceGiatsMarkII := iceGiats // Copies the iceGiants slice (a view of the planets array)
	iceGiats[1] = "Poseidon"
	fmt.Println(planets) // Prints [Mercury Venus Earth Mars Jupiter Saturn Uranus Poseidon]
	fmt.Println(iceGiats, iceGiatsMarkII, ice) // Prints [Uranus Poseidon] [Uranus Poseidon] [Uranus Poseidon]
}