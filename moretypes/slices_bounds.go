package main

import "fmt"

// START OMIT
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4] // HL
	fmt.Println(s[1:4])

	s = s[:2] // HL
	fmt.Println(s[:2])

	s = s[1:] // HL
	fmt.Println(s)
}

// END OMIT
