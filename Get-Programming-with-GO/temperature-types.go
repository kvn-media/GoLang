package main

import (
	"fmt"
)

type celsius float64
type kelvin float64

// kelvinToCelsius converts ºK to ºC
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15) // A type conversion is necessary.
	
}


func main() {
	var k kelvin = 294.0 // The argument must be of type kelvin.
	c := kelvinToCelsius(k)
	fmt.Println(k, "º K is", c,  "º C") // Prints 294º K is 20.850000000000023º C
}