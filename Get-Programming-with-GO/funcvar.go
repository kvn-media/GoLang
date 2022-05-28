package main

import "fmt"

func main() {
	f := func(message string) { // Anonymous func to a variable
		fmt.Println(message)
	}
	f("Go to the party.") // Prints Go to the party
}