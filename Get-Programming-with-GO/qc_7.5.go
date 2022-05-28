package main

import "fmt"

func main() {
	// Add number larger than one
	var red uint8 = 255
	red += 2
	fmt.Println(red) // 1

	var number int8 = 127
	number += 3
	fmt.Println(number) // -126

	// wrap the other way
	red = 0
	red--
	fmt.Println(red) // 255

	number = -128
	number--
	fmt.Println(number) // 127

	// wrapping with a 16-bit unsigned integer
	var green uint16 = 65535
	green++
	fmt.Println(green) // 0
}