package main

import "fmt"

func main() {
	/*
	c := 'a'
	c = c + 3
	fmt.Printf("%c", c) // Prints a

	if c > 'z' {
		c = c - 26
	} // To decipher this Caesar cipher, subtract 3 instead of adding 3. But then you need to account for c < 'a' by adding 26. What a pain.
	*/
	message := "L fdph, L vdz, L frqtxhuhg."
	
	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c -= 3
			if c < 'a' {
				c += 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c -= 3
			if c < 'A' {
				c += 26
			}
		}
		fmt.Printf("%c", c)
	} 
}