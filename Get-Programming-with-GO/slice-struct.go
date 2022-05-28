package main

import "fmt"

func main() {
	/*
	A slice of structures, []struct is a collection of zero or more values (a slice) where each value is based on a structure instead of a primitive type like float64. If a program needed a collection of landing sites for Mars rovers, the way not to do it would be two separate slices for latitudes and longitudes, as shown in the following listing. */
	lats := []float64{-4.5895, -14.5684, -1.9462}
	longs := []float64{137.4417, 175.472636, 354.4734}
	/* This already looks bad, especially in light of the location structure introduced earlier in this lesson. Now imagine more slices being added for altitude and so on. A mistake when editing the previous listing could easily result in data misaligned across slices or even slices of different lengths.*/

	/* A better solution is to create a single slice where each value is a structure. Then each location is a single unit, which you can extend with the name of the landing site or other fields as needed, as shown in the next listing. */
	type location struct {
		name string
		lat  float64
		long float64
	}
	locations := []location{
		{name: "Bradbury Landing", lat: -4.5895, long: 137.4417},
		{name: "Columbia Memorial Station", lat: -14.5684, long: 175.472636},
		{name: "Challenger Memorial Station", lat: -1.9462, long: 354.4734},
	}
	fmt.Println(lats, longs, locations)
}