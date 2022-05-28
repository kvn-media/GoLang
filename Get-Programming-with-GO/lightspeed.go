// How long does it take to get to mars?
package main

import (
	"fmt"
)
// constant and variabel
func main() {
	// const lightspeed = 299792 // km/s
	// var distance = 56000000 // km

	/*
	// shortcut declare multiply variabels at once
	var distance = 56000000
	var speed = 100800
	*/

	// or you can declare as group
	const hoursPerDay = 24
	var (
		distance = 96300000 // km
		speed = 100800 // km/h
	)
	// shortcut
	// var distance, speed = 56000000, 100800

	fmt.Println(distance/speed/hoursPerDay, "days") // print 
	// distance = 401000000
	// fmt.Println(distance/lightspeed, "seconds") // print 1337 seconds

}