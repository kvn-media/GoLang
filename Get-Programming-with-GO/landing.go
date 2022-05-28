package main

import "fmt"

// type location struct {
// 	Name string `json:"name"`
// 	Lat float64 `json:"latitude"`
// 	Long float64 `json:"longitude"`
// }
// locations := []location{
// 	{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
// 	{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
// 	{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
// }
// bytes, err := json.MarshalIndent(locations, "", " ")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	fmt.Println(string(bytes))

// location with a latitude, longitude.
type location struct {
	lat, long float64
}
// coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere.
type coordinate struct {
	d, m, s float64
	h rune
}
// newLocation from latitude, longitude d/m/s coordinates.
func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}
// decimal converts a d/m/s coordinate to decimal degrees.
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func main() {
	spirit := newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	opportunity := newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
	insight := newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0, 'E'})
	fmt.Println("Spirit", spirit)
	fmt.Println("Opportunity", opportunity)
	fmt.Println("Curiosity", curiosity)
	fmt.Println("InSight", insight)
}