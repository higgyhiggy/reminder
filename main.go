package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	id := 99
	name := "newhector"
	fmt.Println("Successfully connected!")
	sqlStatement := `
	INSERT INTO student (id, name)
	VALUES ($1, $2)`
	_, err = db.Exec(sqlStatement, id, name)
	if err != nil {
		panic(err)
	}
       	fmt.Println("this is hector")
}

// fucntion to insert into table

//function to send reminder via email

//fucntion to fetch data from data base

// fucntion to to check if the data of a row in the table is within a certain days away
