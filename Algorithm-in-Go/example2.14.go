package main

import "fmt"

func main() {
	fmt.Println("N = 100, Number of instructions O(n):: ", fun1(100))
	fmt.Println("N = 100, Number of instructions O(n^2):: ", fun2(100))
	fmt.Println("N = 100, Number of instructions O(n^2):: ", fun3(100))
	fmt.Println("N = 100, Number of instructions O(log(n)):: ", fun4(100))
	fmt.Println("N = 100, Number of instructions O(log(n)):: ", fun5(100))
	fmt.Println("N = 100, Number of instructions O(n^3):: ", fun6(100))
	fmt.Println("N = 100, Number of instructions O(n^2):: ", fun7(100))
	fmt.Println("N = 100, Number of instructions O(n^(3/2)):: ", fun8(100))
	fmt.Println("N = 100, Number of instructions O(log(n)):: ", fun9(100))
	fmt.Println("N = 100, Number of instructions O(n^2):: ", fun10(100))
	fmt.Println("N = 100, Number of instructions O(n^3):: ", fun11(100))
	fmt.Println("N = 100, Number of instructions O(n):: ", fun12(100))
	fmt.Println("N = 100, Number of instructions O(n):: ", fun13(100))
}