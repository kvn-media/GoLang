package main

import "fmt"

func main() {
	fmt.Println(fun9)
}

func fun9(n int) int {
	m := 0
	for i := n; i > 0; i /= 2 {
		for j := 0; j < i; j++ {
			m += 1
		}
	}
	return m
}