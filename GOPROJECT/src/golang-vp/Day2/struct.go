package Day2

import "fmt"
// blueprint skematik
type Car struct {
	// nama, color, fuel -> property
	Nama 	string
	Warna  	string
	// Embedded struct / menempelkan struct lain
	Engine Engine

	// nested struct
	// Engine struct
}

type Engine struct {
	TipeMesin 				  string
	IsiSilinder 			  int
	KapasitasTangkiBahanBakar int
}

func SimpleStruct() {
	// inisialisasi sebuah objek
	// carrera sebuah objek
	var Carrera = Car{}
	Carrera.Nama = "Primordial kuning"
	Carrera.Warna = "Jaune"
	Carrera.Engine.TipeMesin = "IL, 4 Cylinder, 16V, DOHC, Dual VVT-i"
	Carrera.Engine.IsiSilinder = 1496
	Carrera.Engine.KapasitasTangkiBahanBakar = 55
	fmt.Println(Carrera)

	var student = []struct {
		name string
		age int
	}{
		{name: "kevin", age: 19},
		{name: "lita", age: 18},
	}

	fmt.Println(student)
}