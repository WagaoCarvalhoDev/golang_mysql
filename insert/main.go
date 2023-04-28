package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:pass-coder@tcp(localhost:3306)/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec("Maria")
	if err != nil {
		panic(err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inseriu %d linha(s) com o ID %d.\n", rows, lastID)

	res, err = stmt.Exec("João")
	if err != nil {
		panic(err)
	}
	rows, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inseriu %d linha(s).\n", rows)

	res, err = stmt.Exec("Pedro")
	if err != nil {
		panic(err)
	}
	lastID, err = res.LastInsertId()
	if err != nil {
		panic(err)
	}
	rows, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inseriu %d linha(s) com o ID %d.\n", rows, lastID)
}
