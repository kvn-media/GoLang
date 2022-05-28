package main

import "fmt"

func main() {
	temperature := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, 33.0,
	}
	frequency := make(map[float64]int)

	for _, t := range temperature { // Iterates over a slice (index, value)
		frequency[t]++
	}
	for t, num := range frequency { // Iterates over a map (key, value)
		fmt.Printf("%+.2f occurs %d times\n", t, num)
	}
}