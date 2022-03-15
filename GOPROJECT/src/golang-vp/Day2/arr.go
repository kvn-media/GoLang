package Day2

import "fmt"

func Arr() {

	// arr 1 dimensi
	var numbers [2]int
	numbers[0] = 1
	numbers[1] = 2
	var numbers2 = [...]int{2, 2, 2, 1, 2, 3}
	for _, v := range numbers2 {
		fmt.Printf("value number %d\n", v)
	}
	// arr 2 dimensi

	// 2 -> row
	// 2 -> colomn
	var numbers3 = [...][2]int{
		{1,2},
		{3,4},
		{4,5},
	}
	fmt.Println(numbers3)
}