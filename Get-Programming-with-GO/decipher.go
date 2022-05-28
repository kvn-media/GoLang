package main

import "fmt"

func main() {
	/* TIP You can only perform operations on values of the same type, but you can convert one type to the other (string, byte, rune).
	   If you recall, modulus gives the remainder of dividing two numbers. For example, 27 % 26 is 1, keeping numbers within the 0–25 range. Be careful with negative numbers, though, as -3 % 26 is still -3.
	*/

	// plainText := "your message goes here"
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	// keyword := "GOLANG"

	for i := 0; i < len(cipherText); i++ {
		c := cipherText[i]
		if c >= 'a' && c <= 'z' {
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
}

/* To be solve */