package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func Connection() *sql.DB {
	var user string = os.Getenv("POSTGRES_USER")
	var dbname string = os.Getenv("POSTGRES_DATABASE")
	var password string = os.Getenv("POSTGRES_PASSWORD")

	connVariables := "user=" + user + " dbname=" + dbname + " password=" + password + " host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connVariables)

	if err != nil {
		panic(err.Error())
	}
	return db
}
