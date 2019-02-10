package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	// START OMIT
	stmt, err := db.Prepare("SELECT id, title, status FROM todos") // HL
	if err != nil {
		log.Fatal("can't prepare query all todos statment", err)
	}

	rows, err := stmt.Query() // HL
	if err != nil {
		log.Fatal("can't query all todos", err)
	}
	for rows.Next() { // HL
		var id int
		var title, status string
		err := rows.Scan(&id, &title, &status) // HL
		if err != nil {
			log.Fatal("can't Scan row into variable", err)
		}
		fmt.Println(id, title, status)
	} // HL

	fmt.Println("query all todos success")
	// END OMIT
}
