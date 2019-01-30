package main

import (
	"fmt"
	"strings"
)

// START OMIT
func WordCount(s string) map[string]int {
	words := strings.Fields(s) // HL
	r := map[string]int{}
	for _, w := range words {
		r[w] = r[w] + 1 // HL
	}
	return r
}

func main() {
	// raw string use `` // HL
	s := `If it looks like a duck, swims like a duck,
		and quacks like a duck, then it probably is a duck.`
	w := WordCount(s) // HL
	fmt.Printf("% #v\n", w)
}

// END OMIT

// someString := "one    two	three four "

// words := strings.Split(someString, " ")

// fmt.Println(words, len(words))

// words2 := strings.Fields(someString)

// fmt.Println(words2, len(words2))
