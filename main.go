package main

import (
	"database/sql"
	"fmt"
	"log"

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
	Check(err)
	defer db.Close()

	err = db.Ping()
	Check(err)

	id := 99
	name := "newhector"
	fmt.Println("Successfully connected!")
	sqlStatement := `
	INSERT INTO student (id, name)
	VALUES ($1, $2)`

	_, err = db.Exec(sqlStatement, id, name)
	Check(err)
	sqlStatement = `SELECT * FROM student WHERE id=$1;`

	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	row := db.QueryRow(sqlStatement, id)
	switch err := row.Scan(&id, &name); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, name)
	default:
		panic(err)
	}

	Check(err)

}

func Check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// fucntion to insert into table
func insert(id int, name string) {

}

//function to send reminder via email

//fucntion to fetch data from data base

// fucntion to to check if the data of a row in the table is within a certain days away
