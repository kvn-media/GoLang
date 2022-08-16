package main

import "fmt"

func main() {
	fmt.Println(fun5)
}

func fun5(n int) int {
	m := 0
	i := n
	for i > 0 {
		m += 1
		i = i / 2
	}
	return m
}