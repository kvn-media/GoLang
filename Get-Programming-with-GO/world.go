package main

import (
	"fmt"
	"math"
)

type location struct {
	lat, long float64
}
// curiosity := location{lat.decimal(), long.decimal()}

type world struct {
	radius float64
}

// newLocation from latitude, longitude d/m/s coordinates.
// func newLocation(lat, long coordinate) location {
// 	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'},
// 	coordinate{137, 26, 30.12, 'E'})
// 	fmt.Println(curiosity)
// 	return location{lat.decimal(), long.decimal()}
// }

var mars = world{radius: 3389.5}
// distance calculation using the Spherical Law of Cosines.
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong) // Uses the world’s radius field
}
// rad converts degrees to radians.
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func main() {
	spirit := location{-14.5684, 175.472636}
	opportunity := location{-1.9462, 354.4734}

	dist := mars.distance(spirit, opportunity) // Uses the distance method on mars
	fmt.Printf("%.2f km\n", dist) // Prints 9669.71 km
}