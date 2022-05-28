package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}
	/* When the Curiosity rover heads east from Bradbury Landing to Yellowknife Bay, the location of Bradbury Landing doesn’t change in real life, nor in the next listing. The curiosity variable is initialized with a copy of the values contained in bradbury, so the values change independently. */
	bradbury := location{-4.5895, 4417}
	curiosity := bradbury

	curiosity.long += 0.0106 // Heads east to Yellowknife Bay
	fmt.Println(bradbury, curiosity) // Prints {-4.5895 137.4417} {-4.5895 137.4523}
}