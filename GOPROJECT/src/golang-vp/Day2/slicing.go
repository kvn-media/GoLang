package Day2

import "fmt"

func SlicingGO() {
	// [2] -> array
	// [...] -> array
	// [] -> slice

	// alamat memorinya 123
	 var names = [3]string{"Ahmad", "Abrar", "Jonathan"}
	// 0 adalah start index, 2 adalah stop sebelum di index ke 2
	// fmt.Println(names[:])
	// names2 -> alamat memori yg sama dengan 123
	// by reference
	// var names2 = names[:]
	// by value
	// var names3sable = names[0:2]
	 fmt.Println(names)
	// fmt.Println(names2)
	// names2[1] = "ibrar"
	// fmt.Println(names)
	// fmt.Println(names2)
	// names = append(names, "joni")

	// fmt.Println(names)
}