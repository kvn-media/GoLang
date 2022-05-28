package main

import "fmt"

/* An explicit pointer to a slice is only useful when modifying the slice itself: the length, capacity, or starting offset. In the following listing, the reclassify function modifies the length of the planets slice. The calling function (main) wouldn’t see this change if reclassify didn’t utilize a pointer. */
func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}
	reclassify(&planets)
	fmt.Println(planets) // Prints [Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]

	/* Instead of mutating the passed slice as in listing 26.18, an arguably cleaner approach is to write the reclassify function to return a new slice. */
}