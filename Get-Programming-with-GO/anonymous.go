package main

import "fmt"

func main() {
	func() { // Declare anonymous function
		fmt.Println("Function anonymous")
	}() // Invokes the function
}