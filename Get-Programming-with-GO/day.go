package main

import (
	"fmt"
	"time"
)

func main() {
	const layout = "Mon, Jan 2, 2006"

	day := time.Now()
	tomorrow := day.Add(24 * time.Hour)

	fmt.Println(day.Format(layout)) // Prints Tue, Nov 10, 2009
	fmt.Println(tomorrow.Format(layout)) // Prints Wed, Nov 11, 2009
}