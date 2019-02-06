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

var todos []Todo

func todosHandler(w http.ResponseWriter, req *http.Request) {
	// START OMIT
	if req.Method == "POST" { // HL
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err)
			return
		}

		t := Todo{}
		err = json.Unmarshal(body, &t)
		if err != nil {
			fmt.Fprintf(w, "error: ", err)
		}

		todos = append(todos, t)    // HL
		fmt.Printf("% #v\n", todos) // HL

		fmt.Fprintf(w, "hello %s created todos", "POST")
		return
	} // HL
	// END OMIT

	fmt.Fprintf(w, "hello %s todos", "GET")
}

func main() {
	http.HandleFunc("/todos", todosHandler)

	log.Fatal(http.ListenAndServe(":1234", nil))
}
