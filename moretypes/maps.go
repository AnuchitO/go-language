package main

import "fmt"

// START OMIT
func main() {
	var m map[string]string     // HL
	m = make(map[string]string) // HL
	m["nong"] = "AnuchitO"
	fmt.Println(m["nong"])
}

// END OMIT
