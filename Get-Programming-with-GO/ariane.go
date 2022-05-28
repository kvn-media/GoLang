package main

import (
	"fmt"
	"math"
)

func main() {
	var bh float64 = 32767
	var h = int16(bh) // To-do: add Rocket Science

	if bh < math.MinInt16 || bh > math.MaxInt16 {
		// handle out of range value
		fmt.Println(bh)
	}
	fmt.Println(h)
}