package main

import (
	"fmt"
	"math"
)

func main() {
	temperature := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	groups := make(map[float64][]float64) // A map with float64 keys and []float64 values

	for _, t := range temperature {
		g := math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
	}
	for g, temperature := range groups {
		fmt.Printf("%v: %v\n", g, temperature)
	}
}