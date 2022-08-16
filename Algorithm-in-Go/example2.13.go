package main

import "fmt"

func main() {
	fmt.Println(fun13)
}

func fun13(n int) int {
	m := 0
	for i := 1; i <= n; i *= 2 {
		for j := 0; j <= i; j++ {
			m += 1
		}
	}
	return m
}