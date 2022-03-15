package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/vica")
	//fmt.Println(err)
	//if err != nil {
	//	panic(err)
	//}
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "", "127.0.0.1", "3306", "vica_book_store")

	db, err := sql.Open(`mysql`, connection)
//	if err != nil {
		log.Fatal(err)
//	}

//	if err := db.Ping(); err != nil {
//		log.Print("err : ", err)
//		_, _ = fmt.Scanln()
//		log.Fatal(err)
//	}

	//var (
	//	id int
	//	name string
	//)

	type book struct {
		Id int
		Name string
	}
		/*
		mhs1 := mahasiswa()
	for rows.Next(){
		err := rows.Scan(&mhs1.id, &mhs1.Name)
	}
		if err != nil{
			log.Fatal(err)
		}
		log.Println(Id, Name)
	}
*/
	// stmt, err := db.Prepare("SELECT * FROM book where name LIKE ?")

	if err != nil {
		log.Fatal(err)
	}

	var data = []book{}
	//book1 := book{}
	//rows, err := stmt.Query("lamp%")

	//for rows.Next() {
	//	err = rows.scan(&book1.Id, &book1.Name)
	//	data = append(data, book1)
	//}

	fmt.Println(data)

	stmt2, err := db.Prepare("INSERT INTO book(name) VALUES (?)")

	stmt2.Exec("Veldora")
	stmt2.Exec("Ramiris")
	stmt2.Exec("Shion")
	stmt2.Exec("Ranga")
	stmt2.Exec("Diablo")
	stmt2.Exec("Cornu")
	stmt2.Exec("Zalario")
	stmt2.Exec("Dino")

	stmtUpdate, err := db.Prepare("UPDATE book SET name = ? where id = ?")
	stmtUpdate.Exec("Veldora", 7)

	stmtDel, err := db.Prepare("DELETE FROM book WHERE name in (?, ?, ?) ")
	stmtDel.Exec("Cornu", "Zalario", "Dino")
//	err = stmt.QueryRow("bobo").Scan(&book1.Id, &book1.Name)
//	if err != nil {
//		log.fatal(err.Error())
//	}
//	data = append(data, book1)
//	fmt.Println(data)

//	book2 := book2{}
//	stmt.QueryRow("lampu merah").Scan(&book2.Id, &book2.Name)
//	data = append(data, book2)

//	fmt.Println(data)

/*
	rows1, err := stmt.Query(1)
	for rows1.Next() {
		err := rows1.Scan(&mhs1.Id, &mhs1.Name)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(mhs1)
	}

	var (
		id int
		name string
	)
*/
//	err = db.QueryRow("SELECT name from book WHERE id = ?", 1).Scan(&name)
//	fmt.Println("nama :", name)

//	defer db.Close()
}
	