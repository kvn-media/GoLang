package main

import "fmt"

func main() {
	message := "assalamualaikum warahmatullahi wabarakatuh"
	for i := 0; i < 42; i++ {
		c := message[i]
		fmt.Printf("%c\n", c)
	}
}