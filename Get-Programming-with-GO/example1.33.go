package main

import "fmt"

func Factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * Factorial(i-1)
}

func main() {
	fmt.Println("factorial 5 is ::", Factorial(5))
}