package main

import (
	"fmt"
	"log"
	"net/http"
)

func todoHandler(w http.ResponseWriter, req *http.Request) { // HL
	method := "GET"
	fmt.Fprintf(w, "hello %s todos", method) // HL
} // HL

func main() {
	http.HandleFunc("/todos", todoHandler) // HL

	log.Fatal(http.ListenAndServe(":1234", nil))
}
