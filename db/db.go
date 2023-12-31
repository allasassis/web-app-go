package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectWithDatabase() *sql.DB {
	connection := "user=postgres dbname=go_store password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}
