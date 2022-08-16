package main

import "fmt"

func main() {
	fmt.Println(fun10)
}

func fun10(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := i; j > 0; j++ {
			m += 1
		}
	}
	return m
}