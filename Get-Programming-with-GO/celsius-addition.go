package main

import "fmt"

func main() {
	type celsius float64
	type fahrenheit float64
	/*
	var c celsius = 20
	var f fahrenheit = 20
	// Invalid operation: mismatched types celsius and fahrenheit, NOTE: Types can’t be mixed
	if c == f {
	}
	c += f
	*/
	const degrees = 20
	var temperature celsius = degrees

	temperature += 10

	var warmUp float64 = 10
	temperature += celsius(warmUp)

	fmt.Println(degrees, temperature, warmUp)
}