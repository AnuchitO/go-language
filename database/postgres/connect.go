package main

import (
	"database/sql" // HL
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // HL
)

func main() {
	//db, err := sql.Open("postgres", "root:password@tcp(127.0.0.1:3306)/tododb")
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL")) // HL
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	fmt.Println("okay")
}
