package main

import "fmt"

type location struct {
	lat, long float64
}

curiosity := location{lat.decimal(), long.decimal()}

// newLocation from latitude, longitude d/m/s coordinates.
func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func main() {
	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'},
		coordinate{137, 26, 30.12, 'E'})
	fmt.Println(curiosity)
}