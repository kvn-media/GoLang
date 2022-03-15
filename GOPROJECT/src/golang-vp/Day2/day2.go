package Day2

import "fmt"

func Operator() {
	num1 := 1
//	num2 := 2
/*
	fmt.Println("result =", num1 + num2)
	fmt.Println("result =", num1 - num2)
	fmt.Println("result =", num1 * num2)
	fmt.Println("result =", num1 / num2)
	// modulus -> sisa bagi
	fmt.Println("result =", num1 % num2)
	//num1++
	num1 = num1 + 1
	fmt.Println(num1)
	num1--
	fmt.Println(num1)
*/
	// Relationl
	/*
	result := num1 == num2
	fmt.Println(result)
	*/
	//var right = false
	//var left = true
	// Logical Operator
	// &&
	// true && true -> true 
	// true && false -> false
	// false && true -> false
	// false && false -> false
	//var leftAndright = left && right
	//fmt.Printf("left && rignt = %t", leftAndright)

	// ||
	// true || true -> true
	// true || false -> true
	// false || false -> false
	// false || true -> true

	//var LeftOrRight = left || right
	//fmt.Printf("left || right =%t", LeftOrRight)

	// ! NOT
	//fmt.Println(!left)
	
	// assignment

	// = assaign sebuah nilai
	// != untuk deklarasi variabel
	// +=
	num1 += 2


	// bitwise operator
	// 2^0 2^1 2^2 2^3
	bit1 := 4 // 100
	bit2 := 6 // 110 -> 11000

	// & 100 & 110 = 100
	fmt.Println(bit1 & bit2)
	// | 100 | 110 = 120
	fmt.Println(bit1 | bit2)
	// ^ 100 ^ 110 = 010
	fmt.Println(bit1 ^ bit2)
	// &^ -> AND NOT 100 ^ 100 = 010
	fmt.Println(bit1 &^ bit2)
	// << left shift 100 << 2
	fmt.Println(bit2<<2)
}