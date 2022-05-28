package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64
// celsius converts ºF to ºC
func (f fahrenheit) celsius() celsius {
return celsius((f - 32.0) * 5.0 / 9.0)
} // This creates a nice symmetry, where every type of temperature can have a celsius method to convert to Celsius.

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func main() {
	var c celsius = 127.0
	k := celsiusToKelvin(c)
	fmt.Print(c, "º C is ", k, "º K")
}

/* The syntax to use a method is different than calling a function:
   var k kelvin = 294.0
   var c celsius
   c = kelvinToCelsius(k)
   c = k.celsius()
*/