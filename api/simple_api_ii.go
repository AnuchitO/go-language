package main

import (
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	resp := []byte(`{"name": "anuchit"}`)
	w.Write(resp)
	io.WriteString(w, "Hello")
}

func main() {
	http.HandleFunc("/", helloHandler)

	err := http.ListenAndServe(":1234", nil) // HL
	log.Fatal(err)                           // HL

	// normally make it one line
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
