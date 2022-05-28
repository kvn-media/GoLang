package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Cómo estás?"

	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}

	/*
	for _, c := range question {
		fmt.Printf("%c\n", c) // Prints ¿ C ó m o e s t á s ?
	}
	*/

	fmt.Println(len(question), "bytes") // Prints 15 bytes
	fmt.Println(utf8.RuneCountInString(question), "runes") // Prints 12 runes

	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes", c, size) // Prints First rune: ¿2 bytes
}