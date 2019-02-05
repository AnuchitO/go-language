package main

import (
	"fmt"
	"log"
	"net/http"
)

func todosHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" { // HL
		fmt.Fprintf(w, "hello %s created todos", "POST")
		return // HL
	}

	fmt.Fprintf(w, "hello %s todos", "GET")
}

func main() {
	http.HandleFunc("/todos", todosHandler)

	log.Fatal(http.ListenAndServe(":1234", nil))
}
