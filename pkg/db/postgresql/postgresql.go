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

	//a := db.QueryRow("SELECT text FROM public.comments LIMIT 1")
	//var k string
	//a.Scan(&k)
	//fmt.Println(k)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to PostgreSQL")

	return db
}
