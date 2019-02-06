package main

import (
	"encoding/json"
	"fmt"
)

// START OMIT
type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func main() {
	t := Todo{ // HL
		ID:     "1",
		Title:  "pay credit card",
		Status: "completed",
	} // HL

	b, err := json.Marshal(t) // HL
	fmt.Printf("%T => %v \n %s \n", b, b, b)
	fmt.Println(err)
}

// END OMIT
