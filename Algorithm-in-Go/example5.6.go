package main

import (
	"fmt"
	"sort"
)

func printRepeating2(data []int) {
	size := len(data)
	sort.Ints(data) // Sort(data,size)
	fmt.Println("Repeating elements are: ")

	for i := 1; i < size; i++ {
		if data[i] == data[i-1] {
			fmt.Println("", data[i])
		}
	}
	fmt.Println()
}

func main() {

}