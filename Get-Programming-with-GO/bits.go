package main

import "fmt"

func main() {
	var green uint8 = 3
	fmt.Printf("%08b\n", green) // 00000011
	green++
	fmt.Printf("%08b\n", green) // 00000100
}