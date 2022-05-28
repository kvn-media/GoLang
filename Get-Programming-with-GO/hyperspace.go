package main

import (
	"fmt"
	"strings"
)
type StringSlice []string

func (p StringSlice) Sort()

func hyperspace(worlds []string) { // This argument is a slice, not an array.
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func main() {
	planets := []string{" Venus		", "Earth	", "Mars"} // Planet surrounded by space
	hyperspace(planets)

	fmt.Println(strings.Join(planets, "")) // Prints VenusEarthMars
	/* Both worlds and planets are slices, and though worlds is a copy, they both point to the same underlying array. */
}

/*
func (p StringSlice) type SortBy []Type

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }
*/