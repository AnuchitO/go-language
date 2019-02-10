package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// START OMIT
func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO todos (title, status) values ($1, $2)") // HL
	if err != nil {
		log.Fatal("can't prepare statment insert", err)
	}

	if _, err := stmt.Exec("buy bmw", "active"); err != nil { // HL
		log.Fatal("execute insert error", err)
	} // HL

	fmt.Println("insert todo success")
}

// END OMIT
