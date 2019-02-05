package main

import (
	"net/http" // HL
)

func helloHandler(w http.ResponseWriter, req *http.Request) { // HL
	resp := []byte(`{"name": "anuchit"}`)
	w.Write(resp)
} // HL

func main() {
	http.HandleFunc("/", helloHandler) // HL

	http.ListenAndServe(":1234", nil) // HL
}
