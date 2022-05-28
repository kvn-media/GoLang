// My weight loss program.
package main

import (
	"fmt"
)
// main is the function where it all begins
func main() {
	// Perform calculation
	fmt.Println("My weight on the surface of mars is ")
	fmt.Println(149.0 * 0.3783) // print 56.3667
	fmt.Println(" lbs, and I would be ")
	fmt.Println(41 * 365 / 687) // print 21
	fmt.Println(" years old.")

	// Formatted Print
	fmt.Printf("My weight on the surface of mars is %v lbs,", 149.0*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 41*365/687)
}