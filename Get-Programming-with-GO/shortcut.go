package main

import "fmt"

func main() {
	var weight = 149.0
	weight = weight * 0.3783
	weight *= 0.3783
	fmt.Println(weight)

	// increment operator ++ , decrement operator count-- shorten other operations like price /= 2 in the same way.
	var age = 41
	age = age + 1
	age += 1
	age++
}