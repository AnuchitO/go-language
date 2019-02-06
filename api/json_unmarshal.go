package main

// START OMIT
import (
	"encoding/json" // HL
	"fmt"
)

type Todo struct {
	ID     string `json:"id"`     // HL
	Title  string `json:"title"`  // HL
	Status string `json:"status"` // HL
}

func main() {
	data := []byte(`{
			"id": "1",
			"title": "pay credit card",
			"status": "active"
		}`)

	var t Todo
	err := json.Unmarshal(data, &t) // HL

	fmt.Printf("% #v\n", t)
	fmt.Println(err)
}

// END OMIT
