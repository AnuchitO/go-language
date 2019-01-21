package main

import "fmt"

// START OMIT
func main() {
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	a := names[0:2] // HL
	b := names[1:3] // HL
	fmt.Println(a, b)

	b[0] = "XXX" // HL
	fmt.Println(a, b)
	fmt.Println(names) // HL
}

// END OMIT
