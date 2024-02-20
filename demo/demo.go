package main

import (
	"database/sql"
	"log"
)

func main() {

	db, err := sql.Open("sqlite3", "gee.db")
	if err != nil {
		log.Println(err)
	}
	defer func() {
		_ = db.Close()
	}()
	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
	_, _ = db.Exec("CREATE TABLE User(Name text);")

	result, err := db.Exec("INSERT INTO User(Name) values (?), (?)", "Tom", "Sam")

	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)

	}
	row := db.QueryRow("SELECT Name FROM User LIMIT 1")

	var name string

	if err := row.Scan(&name); err == nil {
		log.Println(name)
	}
}
