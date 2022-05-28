package main

import "fmt"

// The number of elements that are visible through a slice determines its length. If a slice has an underlying array that is larger, the slice may still have capacity to grow. The following listing declares a function to print out the length and capacity of a slice.
func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice),
	cap(slice), slice)
}

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dump("dwarfs", dwarfs) // Prints dwarfs: length 5, capacity 5 [Ceres Pluto Haumea Makemake Eris]
	dump("dwarfs[1:2]", dwarfs[1:2]) // Prints dwarfs[1:2]: length 1, capacity 4 [Pluto]
	// The slice created by dwarfs[1:2] has a length of 1, but the capacity to hold 4 elements.
	// Note: Pluto Haumea Makemake Eris provide a capacity of 4 even though the length is 1.
}