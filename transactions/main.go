package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:pass-coder@tcp(localhost:3306)/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	stmt, err := tx.Prepare("INSERT INTO users(id, name) VALUES(?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	stmt.Exec(3000, "Bia")
	stmt.Exec(3001, "Carlos")

	//_, err = stmt.Exec(1, "Thiago")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
