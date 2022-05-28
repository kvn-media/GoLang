package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15, // Composite literals are key-value pairs for maps.
		"Mars":  -65,
	}
	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %vº C.\n", temp) // Prints On average the Earth is 15º C.

	temperature["Earth"] = 16 // A little climate change
	temperature["Venus"] = 464

	fmt.Println(temperature) // Prints map[Venus:464 Earth:16 Mars:-65]

	moon := temperature["Moon"]
	fmt.Println(moon) // Prints 0

	if moon, ok := temperature["Moon"]; ok { // The comma, ok syntax
		fmt.Printf("On average the moon is %vº C.\n", moon)
	} else {
		fmt.Println("Where is the moon?") // Prints Where is the moon?
	}
}