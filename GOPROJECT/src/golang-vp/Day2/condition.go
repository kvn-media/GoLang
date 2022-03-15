package Day2

import "fmt"

func Condition() {
	angka1 := 2
	angka2 := 2
	angka3 := 3
/*
	// if else statement
	if angka1 > angka2 {
		fmt.Println("gokurosan")
	} else {
		fmt.Println("semimasen")
	}

	// else if statement
	if angka1 < angka2 {
		fmt.Println("gokurosan")
	} else if angka1 == angka2 && angka1 == angka3 {
		fmt.Println("nilai sama")
	} else {
		fmt.Println("semimasen")
	}
*/
	// nested if
	if angka1 == angka2 {
		if angka1 != angka3 {
			fmt.Println("yes kondisi terpenuhi")
		}
	}

	day := "senin"
	// switch
	switch day {
	case "senin":
		fmt.Println("MTK")
	case "selasa":
		fmt.Println("Agama")
		default:
			fmt.Println("eskul")	
	}
}