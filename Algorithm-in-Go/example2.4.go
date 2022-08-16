package main

import "fmt"

func main() {
	fmt.Println(fun4)
}

func fun4(n int) int {
	m := 0
	i := 1
	for i < n {
		m += 1
		i = i * 2
	}
	return m
}