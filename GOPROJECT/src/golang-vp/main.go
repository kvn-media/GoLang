package main

import (
	"command-go/Day2"
	"fmt"
)

/*
var NAMA string = "ENIGMA VICA PREMIUM GOLANG"

// {} -> block scope
func main() {
	fmt.Println("NAMA", NAMA)
	// cara deklarasi variable di golang
	// cara pengisian variable dari kanan ke kiri
	// = itu digunakan untuk assign sebuah nilai / value
	// string
	// cara 1 deklarasi variable
	var wadah1 string = "Test"
	// cara 2
	wadah2 := 'A'

	isActive := true
	// ini digunakan untuk mencetak
	fmt.Println(wadah1)
	fmt.Println(wadah2)
	fmt.Println(isActive)


	// tipe data

	// boolean -> true or false
	// uint
	// var angka uint8 = 255
	// int8 -128 < - > 127
	// var sampleInt8 int8 = -127
	// int16 -32768 < - > 32767
	// int32 -2147483648 < - > 2147483647
	// int64
	// int sama int32 sama int64
	// rune sama dengan int32

	// float32
	// var desc float32

	// zero value
	// tidak memiliki value, dia akan membuat nilai default
	var number1 int = 20
	number1 = 30
	var isOn bool
	fmt.Println("ini number", number1)
	fmt.Println("ini isIn", isOn)

	const phi = 3.14
	greet()
	// phi = 2.77

}

func greet() {
	fmt.Println("Hello", NAMA)
}
*/

func main(){
//	Day2.Operator()
//	Day2.Condition()
//	Day2.Loop()
//	Day2.Arr()
//	Day2.SlicingGO()
//	Day2.Maps()
//	Day2.Greet("world", 10)
//	result, err := Day2.Sum(0,2)
//	if result == 0 {
//		fmt.Println(err)
//	}
//	fmt.Println(result)
//	Day2.SetNames("test", "kevin")
//	greeting := Day2.VariabelFunc()
//	greeting()
//	Day2.FuncForParam(greetings)
//	Day2.VariabelFunc()()
//	Day2.Pointer()
//	Day2.SimpleStruct()
	rect := Day2.Rectangle{}
	// gak proper
	// react.length = 10
	// merubah nilai width
	rect.SetWidth(10)
	rect.SetLength(2)
	fmt.Println(rect.GetWidth())
	fmt.Println(rect.GetArea())

	var shape Day2.Shape

	shape = Day2.Rectangles{Width: 10.0, Length: 2.0}
	fmt.Println(shape.GetArea())

	shape = &Day2.Circle{}
	shape.(*Day2.Circle).SetRadius(10.0)
	fmt.Println(shape.GetArea())
	Day2.EmptyInterFace()
}


//func Calculate(x int) (result int) {
//	result = x + 2
//	return result
//}


