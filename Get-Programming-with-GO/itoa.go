package main

import (
	"fmt"
	"strconv"
)

func main() {
	// countdown := 10
	// str := "Launch in T minus " + strconv.Itoa(countdown) + " seconds."
	// fmt.Println(str) // Prints Launch in T minus 10 seconds

	// Another way to convert number to a string is to use Sprintf, a cousin of Printf that returns a string rahter than displaying it:
	countdown := 9
	str := fmt.Sprintf("Launch in T minus %v seconds.", countdown)
	fmt.Println(str) // Prints Launch in T minus 9 seconds.

	countdown, err := strconv.Atoi("10")
	if err != nil {
 	// oh no, something went wrong
	}
	fmt.Println(countdown) // 10
	/* To go the other way, the strconv package provides the Atoi function (ASCII to integer). Because a string may contain gibberish or a number that’s too big, the Atoi function may return an error: */
}