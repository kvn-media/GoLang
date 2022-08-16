package main

import "fmt"

func printRepeating3(data []int) {
	s := make(Set)
	size := len(data)
	fmt.Println("Repeating elements are:")
	for i := 0; i < size; i++ {
		if s.Find(data[i]) {
			fmt.Println("", data[i])
		} else {
			s.Add(data[i])
		}
	}
	fmt.Println()
}

func main() {

}