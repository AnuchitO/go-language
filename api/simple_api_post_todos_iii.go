package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func todosHandler(w http.ResponseWriter, req *http.Request) {
	// START OMIT
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err)
			return
		}

		t := Todo{}                    // HL
		err = json.Unmarshal(body, &t) // HL
		if err != nil {                // HL
			fmt.Fprintf(w, "error: ", err)
		} // HL
		fmt.Printf("body : % #v\n", t) // HL

		fmt.Fprintf(w, "hello %s created todos", "POST")
		return
	}
	// END OMIT

	fmt.Fprintf(w, "hello %s todos", "GET")
}

func main() {
	http.HandleFunc("/todos", todosHandler)

	log.Fatal(http.ListenAndServe(":1234", nil))
}
