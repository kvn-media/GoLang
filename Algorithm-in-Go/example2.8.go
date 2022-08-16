package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(fun8)
}

func fun8(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		sq := math.Sqrt(float64(n))
		for j := 0; j < int(sq); j++ {
			m += 1
		}
	}
	return m
}