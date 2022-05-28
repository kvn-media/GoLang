package main

import "fmt"
// Iterating through an array with range
func main() {
	dwarfs := [5]string{"Ceres", "Pluto", "Heumea", "Makemake", "Eris"}

	for i, dwarfs := range dwarfs {
		fmt.Println(i, dwarfs)
	}
}