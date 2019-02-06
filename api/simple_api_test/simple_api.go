package simple

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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

	// w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "hello %s todos", "GET")
}
