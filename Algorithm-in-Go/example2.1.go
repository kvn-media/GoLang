package main

import "fmt"

func main() {
	fmt.Println(fun1)
}

func fun1(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		m += 1
	}
	return m
}