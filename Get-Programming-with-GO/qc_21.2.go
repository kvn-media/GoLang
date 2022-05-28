package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}
	var curiosity location
	curiosity.lat = -4.5895
	curiosity.long = 137.4417
	fmt.Println(curiosity)
}