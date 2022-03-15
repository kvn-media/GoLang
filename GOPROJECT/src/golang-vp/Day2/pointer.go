package Day2

import "fmt"

func Pointer() {
	var angka int = 1
//	var angka1 *int
//	angka1 = &angka
//	fmt.Println(&angka)
//	fmt.Println(*angka1)
	fmt.Println("Angka sebelum perubahan", angka)
	ChangeNumber(&angka, 10)
	fmt.Println("Angka setelah perubahan", angka)
}

func ChangeNumber(number *int, value int) {
	*number = value
	fmt.Println(number)
}