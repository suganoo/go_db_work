package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=postgres password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func main() {
	rows, _ := Db.Query("SELECT * FROM posts ORDER BY id")
	for rows.Next() {
		var id int
		var content string
		var author string
		rows.Scan(&id, &content, &author)
		fmt.Println(id, content, author)
	}
}
