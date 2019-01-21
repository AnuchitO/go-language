package main

import "fmt"

// START OMIT
func main() {
	m := make(map[string]int)

	m["Answer"] = 42 // HL
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48 // HL
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer") // HL
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] // HL
	fmt.Println("The value:", v, "Present?", ok)
}

// END OMIT
