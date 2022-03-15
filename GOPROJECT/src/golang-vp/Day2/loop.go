package Day2

import "fmt"

func Loop() {
/*	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}	
	}

	// for while
	num1 := 1
	for num1 < 10{
		fmt.Println(num1)
		num1 ++
	}
*/
	// for ever
	num1 := 10
/*	for {
		fmt.Println("")
		if num1 == 1 {
			fmt.Println("")
			break
		}
	}
*/
	for num1 < 18 {
		if num1 == 15 {
			continue
		}
		fmt.Println("value", num1)
		num1++
	}
}