package main

import "fmt"

func SumArray(data []int) int {
	size := len(data)
	total := 0
	for index := 0; index < size; index++ {
		total = total + data[index]
	}
	return total
}

func main() {
	fmt.Println(SumArray)
}
