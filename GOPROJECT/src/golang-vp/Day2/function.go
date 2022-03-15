package Day2

import (
	"errors"
	"fmt"
)

// value adalah sebuah parameter -> untuk
func Greet(value string, number int) {
	fmt.Println("hello " + value, number)
}

// fungsi return
func Sum(angka1, angka2 int) (int, error) {
	if angka1 == 0 || angka2 == 0 {
		return 0, errors.New("number cannot zero")
	}		
	result := angka1 + angka2
	return result, nil
}

func SetNames(name string, names ...string) {
	fmt.Println(name, names)
}

func VariabelFunc() func() {
	greeting := func () {
		fmt.Println("Hello")
	}
	 return greeting
}

func FuncForParam(greet func()) {
	greet()
}