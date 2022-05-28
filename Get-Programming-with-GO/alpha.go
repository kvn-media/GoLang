package main

import "fmt"

func main() {
	const lightspeed = 299792 // km/s
	const secondsPerDay = 86400

	var distance int64 = 41.3e12
	fmt.Println("Alpha Centauri is", distance, "km away.") // Prints Alpha Centauri is 41300000000000 km away

	days := distance / lightspeed / secondsPerDay
	fmt.Println("That is", days, "days of travel at light speed.") // Prints That is 1594 days of travel at light speed
}