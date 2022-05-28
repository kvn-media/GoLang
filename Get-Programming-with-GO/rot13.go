package main

import "fmt"

func main() {
	message := "uv vagreangvbany fcnpr fgngvba"

	for i := 0; i < len(message); i++ { // Iterates through each ASCII character
		c := message[i]
		if c >= 'a' && c <= 'z' { // Leaves spaces and punctuation as they are
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
		/*
		Note that the ROT13 implementation in the previous listing is only intended for ASCII characters (bytes). It will get confused by a message written in Spanish or Russian. The next section looks at a solution for this issue.
		*/
	}
}