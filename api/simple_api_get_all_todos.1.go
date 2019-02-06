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

	if req.Method == "POST" {
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

		todos = append(todos, t)
		fmt.Printf("% #v\n", todos)

		fmt.Fprintf(w, "hello %s created todos", "POST")
		return
	}

	// START OMIT
	if req.Method == "GET" {
		b, err := json.Marshal(todos) // HL
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error: ", err)
			return
		}

		w.Header().Set("Content-Type", "application/json") // HL
		w.Write(b)                                         // HL
	}
	// END OMIT
}

func main() {
	http.HandleFunc("/todos", todosHandler)

	log.Fatal(http.ListenAndServe(":1234", nil))
}
