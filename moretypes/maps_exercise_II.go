package main

import "fmt"

func WordCount(s string) map[string]int { // HL
	return map[string]int{"x": 1}
} // HL

func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."
	w := WordCount(s) // HL
	fmt.Println(w)
}
