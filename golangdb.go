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

type Member struct {
	Id         int
	FirstName  string
	LastName   string
	Email      string
	AccessPrev bool
}

func main() {
	fmt.Println("----- Query 1 -----")
	rows, err := Db.Query("SELECT * FROM members ORDER BY id")
	if err != nil {
		return
	}
	for rows.Next() {
		m := Member{}
		rows.Scan(&m.Id, &m.FirstName, &m.LastName, &m.Email, &m.AccessPrev)
		fmt.Println(m)
	}
	fmt.Println()

	fmt.Println("----- Query 2 -----")
	rows, err = Db.Query("SELECT * FROM members WHERE accessprev = $1 ORDER BY id", "false")
	if err != nil {
		return
	}
	for rows.Next() {
		m := Member{}
		rows.Scan(&m.Id, &m.FirstName, &m.LastName, &m.Email, &m.AccessPrev)
		fmt.Println(m)
	}
	fmt.Println()

	fmt.Println("----- QueryRow -----")
	m := Member{}
	err = Db.QueryRow("SELECT * FROM members WHERE accessprev = $1 ORDER BY id", "true").Scan(&m.Id, &m.FirstName, &m.LastName, &m.Email, &m.AccessPrev)
	if err != nil {
		return
	}
	fmt.Println(m)
	fmt.Println()

	fmt.Println("----- Prepare -----")
	statement := "SELECT * FROM members WHERE accessprev = $1 ORDER BY id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err = stmt.Query("false")
	for rows.Next() {
		m := Member{}
		rows.Scan(&m.Id, &m.FirstName, &m.LastName, &m.Email, &m.AccessPrev)
		fmt.Println(m)
	}
	fmt.Println()

	fmt.Println("----- Named -----")
	rows, err = Db.Query(`SELECT * FROM members WHERE id = :tanaka`, sql.Named("tanaka", 1))
	fmt.Println(err)
	if err != nil {
		return
	}
	for rows.Next() {
		m := Member{}
		rows.Scan(&m.Id, &m.FirstName, &m.LastName, &m.Email, &m.AccessPrev)
		fmt.Println(m)
	}
	fmt.Println()

}
