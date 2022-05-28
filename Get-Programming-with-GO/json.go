package main

import (
	"encoding/json"
	"fmt"
	"os"
)
/* JavaScript Object Notation, or JSON (json.org), is a standard data format popularized by Douglas Crockford. It’s based on a subset of the JavaScript language but it’s widely supported in other programming languages. JSON is commonly used for web APIs (Application Programming Interfaces), including the MAAS API (github.com/ingenology/mars_weather_api) that provides weather data from the Curiosity rover. 

The Marshal function from the json package is used in listing 21.9 to encode the data in location into JSON format. Marshal returns the JSON data as bytes, which can be sent over the wire or converted to a string for display. It may also return an error, a topic that’s covered in lesson 28. */
func main() {
	type location struct {
		Lat, Long float64 // Fields must begin with an uppercase letter.
	}
	curiosity := location{-4.5895, 137.4417}

	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Println(string(bytes)) // Prints {"Lat":-4.5895, "Long":137.4417}
}
// exitOnError prints any error and exits
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
// Notice that the JSON keys match the field names of the location structure. For this to work, the json package requires fields to be exported. If Lat and Long began with a lowercase letter, the output would be {}.