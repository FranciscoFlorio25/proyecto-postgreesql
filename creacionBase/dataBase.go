package creacionBase

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func createDataBase() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`drop database if exists tarjetascredito`)
	_, err = db.Exec(`
		create database tarjetascredito;
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func CrearBase() {

	createDataBase()

	fmt.Printf("Bases de datos creada\n")
}
