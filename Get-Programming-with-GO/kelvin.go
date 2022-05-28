package main

import "fmt"

func kelvinToCelsius(k float64) float64 { // Declares a function that accepts one parameter and returns one result
	k -= 273.15
	return k
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin) // Calls the function passing kelvin as the first argument
	fmt.Println(kelvin,  "º K is ", celsius, "º C") // Prints 294º K is 20.850000000000023º C
}