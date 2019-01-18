package main

import "fmt"

// START OMIT
func main() {
	names := []string{"game", "bank", "samui", "dog", "jiew"}
	for _, name := range names { // HL
		fmt.Println(name)
	} // HL

	for { // HL
		fmt.Println("print forever.")
	} // HL
}

// END OMIT
