package main

import "fmt"

func printRepeating(data []int) {
	size := len(data)
	fmt.Println("Repeating elements are:")
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				fmt.Println("", data[i])
			}
		}
	}
	fmt.Println()
}

func main() {

}