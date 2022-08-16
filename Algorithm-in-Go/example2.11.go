package main

import "fmt"

func main() {
	fmt.Println(fun11)
}

func fun11(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			for k := j + 1; k < n; k++ {
				m += 1
			}
		}
	}
	return m
}