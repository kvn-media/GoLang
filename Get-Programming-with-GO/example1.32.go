package main

import "fmt"

// System stack and Method calls
func f2() {
	fmt.Println("fun2 line1")
}

func f1() {
	fmt.Println("fun1 line 1")
	f2()
	fmt.Println("fun 1 line 2")
}

func main() {
	fmt.Println("main line 1")
	f1()
	fmt.Println("main line 2")
}