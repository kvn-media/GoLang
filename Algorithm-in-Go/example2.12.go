package main

import "fmt"

func main() {
	fmt.Println(fun12)
}

func fun12(n int) int {
	j := 0
	m := 0
	for i := 0; i < n; i++ {
		for ; j < n; j++ {
			m += 1
		}
	}
	return m
}