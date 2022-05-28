package main
import (
	"fmt"
	"math/rand"
)

func main() {
	/*
	piggyBank := 0.0
	for piggyBank < 20.00 {
		switch rand.Intn(3) {
		case 0:
			piggyBank += 0.05
		case 1:
			piggyBank += 0.10
		case 2:
			piggyBank += 0.25
	}
	fmt.Printf("$%5.2f\n", piggyBank)
	
	*/
	
	piggyBank := 0
	for piggyBank < 2000 {
		switch rand.Intn(3) {
		case 0:
			piggyBank += 5
		case 1:
			piggyBank += 10
		case 2:
			piggyBank += 25
	}
	dollars := piggyBank / 100
	cents := piggyBank % 100
	fmt.Printf("$%d.%02d\n", dollars, cents)
	}
}	