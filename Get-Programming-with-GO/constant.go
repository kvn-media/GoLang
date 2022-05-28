package main

import "fmt"

func main() {
	const distance = 24000000000000000000
	const lightspeed = 299792
	const secondsPerDay = 86400

	const days = distance / lightspeed / secondsPerDay
	fmt.Println("Andromeda Galaxy is", days, "light days away") // Prints Andromeda Galaxy is 926568346 light days away.
}