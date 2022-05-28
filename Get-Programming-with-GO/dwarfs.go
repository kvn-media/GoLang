package main

import "fmt"
// Looping through an array
func main() {
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	for i := 0; i < len(dwarfs); i++ {
		dwarfs := dwarfs[i]
		fmt.Println(i, dwarfs)
	}
}