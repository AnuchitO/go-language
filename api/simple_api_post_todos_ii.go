package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// START OMIT
func todosHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body) // HL
		if err != nil {                       // HL
			fmt.Fprintf(w, "error : %v", err)
			return // HL
		}

		fmt.Printf("body : %s\n", body) // HL

		fmt.Fprintf(w, "hello %s created todos", "POST")
		return
	}

	fmt.Fprintf(w, "hello %s todos", "GET")
}

// END OMIT
func main() {
	http.HandleFunc("/todos", todosHandler)

	log.Fatal(http.ListenAndServe(":1234", nil))
}
