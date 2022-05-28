package main

import "fmt"

func main() {
	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"} // Length 5, capacity
	dwarfs2 := append(dwarfs1, "Orcus")                                 // Length 6, capacity 10
	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")            // Length 9, capacity 10
	fmt.Println(dwarfs1, dwarfs2, dwarfs3)
}