package main

import "fmt"

func main() {
	fmt.Println(fun3)
}

func fun3(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			m += 1
		}
	}
	return m
}