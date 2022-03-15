package Day2

import "fmt"

func Maps() {
	var data = map[int]string{}
	data[1] = "abi"
	data[2] = "abu"
	// _, isExist := data[3]
	delete(data, 1)
	fmt.Println(data)

	var data2 = map[string]string{
		"nama" : "enigma",
		"lokasi" : "jakarta",
	}
	fmt.Println(data2)

	for key, value := range data{
		fmt.Println("Key :", key, "Value", value)
	}

	mahasiswa := make(map[string]string)
	mahasiswa["nama"] = "Horikita"
	mahasiswa["kelas"] = "D-1"
	fmt.Println(mahasiswa)

	dataMahasiswa := []map[string]string{
		{"nama" : "Horikita", "kelas" : "D1"},
		{"nama" : "Ayanakoji", "kelas" : "D-1"},
	}
	fmt.Println(dataMahasiswa)

	mhsNew := map[string]string{"nama" : "tika", "kelas" : "B501"}

	dataMahasiswa = append(dataMahasiswa, mhsNew)
	fmt.Println(dataMahasiswa)
}