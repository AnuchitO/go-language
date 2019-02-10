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

	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1") // HL
	if err != nil {
		log.Fatal("can'tprepare query one row statment", err)
	}

	rowId := 1
	row := stmt.QueryRow(rowId) // HL
	var id int
	var title, status string

	err = row.Scan(&id, &title, &status) // HL
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}

	fmt.Println("one row", id, title, status)
}

// END OMIT
