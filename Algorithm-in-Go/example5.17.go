package main

import "fmt"

func findPair(data []int, value int) bool {
	size := len(data)
	ret := false
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if (data[i] + data[j]) == value {
				fmt.Println("The pair is: ", data[i], "", data[j])
				ret = true
			}
		}
	}
	return ret
}

func main() {

}