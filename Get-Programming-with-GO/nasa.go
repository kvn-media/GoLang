package main

import "fmt"

func main() {
	var administrator *string

	scolese := "Christopher J. Scolese"
	administrator = &scolese
	fmt.Println(*administrator) // Prints Christopher J. Scolese

	bolden := "Charles Frank Bolden Jr."
	administrator = &bolden
	fmt.Println(*administrator) // Prints Charles Frank Bolden Jr

	*administrator = "Maj. Gen. Charles Frank Bolden Jr."
	fmt.Println(bolden)

	major := administrator
	*major = "Major General Charles Frank Bolden Jr."
	fmt.Println(bolden)

	fmt.Println(administrator == major) // Prints True

	lightfoot := "Robert M. Lightfoot Jr."
	administrator = &lightfoot
	fmt.Println(administrator == major) // Prints False

	charles := *major
	*major = "Charles Bolden"
	fmt.Println(charles) // Prints Major General Charles Frank Bolden Jr.
	fmt.Println(bolden) // Prints Charles Bolden

	charles = "Charles Bolden"
	fmt.Println(charles == bolden) // Prints true
	fmt.Println(&charles == &bolden) // Prints false
}