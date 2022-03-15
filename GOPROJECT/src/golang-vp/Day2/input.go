package Day2

import (
	"bufio"
	"fmt"
	"os"
)

func Input() {
	var angka int
	fmt.Print("Masukan angka : ")
	fmt.Scan(&angka)
	fmt.Println("Angka yang dimasukan : ", angka)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukan kalimat : ")
	scanner.Scan()
	fmt.Println("Kalimat yang dimasukan : '" + scanner.Text() + "'")
}