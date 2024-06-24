package database

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "a6d3fdd4"
	dbname   = "ozon_tech"
)

var connectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

func InitBD() *sql.DB {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to PostgreSQL")

	return db
}
