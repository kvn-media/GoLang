package main

import "fmt"

func main() {
	c := 'a'
	c = c + 3
	fmt.Printf("%c", c) // Prints a

	if c > 'z' {
		c = c - 26
	} // To decipher this Caesar cipher, subtract 3 instead of adding 3. But then you need to account for c < 'a' by adding 26. What a pain. 
}