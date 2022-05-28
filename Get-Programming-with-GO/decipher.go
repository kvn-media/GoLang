package main

import "fmt"

func main() {
	/* TIP You can only perform operations on values of the same type, but you can convert one type to the other (string, byte, rune).
	   If you recall, modulus gives the remainder of dividing two numbers. For example, 27 % 26 is 1, keeping numbers within the 0–25 range. Be careful with negative numbers, though, as -3 % 26 is still -3.
	*/
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	message := ""
	keyIndex := 0

	for i := 0; i < len(cipherText); i++ {
		// A=0, B=1, ... Z=25
		c := cipherText[i] - 'A'
		k := keyword[keyIndex] - 'A'

		// cipher letter - key letter
		c = (c-k+26)%26 + 'A'
		message += string(c)

		// increment keyIndex
		keyIndex++
		keyIndex %= len(keyword)
	}
	fmt.Println(message)
}	
/* To be solve */