package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int { // HL
	words := strings.Fields(s)
	r := map[string]int{}
	for _, w := range words {
		r[w] = r[w] + 1
	}
	return r
} // HL

func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."
	w := WordCount(s) // HL
	fmt.Printf("% #v\n", w)
}
