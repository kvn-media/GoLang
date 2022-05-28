package main

import "fmt"

func main() {
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}
	planetsMarkII := planets
	planets["Earth"] = "whoops"

	fmt.Println(planets) // Prints map[Earth:whoops Mars:Sector ZZ9]
	fmt.Println(planetsMarkII) // // Prints map[Earth:whoops Mars:Sector ZZ9]

	delete(planets, "Earth") // Removes Earth from the map
	fmt.Println(planetsMarkII) // Prints map[Mars:Sector ZZ9]
}