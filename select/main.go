package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id   int
	name string
}

func main() {
	db, err := sql.Open("mysql", "root:pass-coder@tcp(localhost:3306)/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM users WHERE id > ?", 3)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var u user
		rows.Scan(&u.id, &u.name)
		fmt.Println(u)
	}
}
