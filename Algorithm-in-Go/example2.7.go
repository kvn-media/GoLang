package main

import "fmt"

func main() {
	fmt.Println(fun7)
}

func fun7(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m += 1
		}
	}
	return m
}