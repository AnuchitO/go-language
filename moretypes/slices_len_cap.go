package main

import "fmt"

// START OMIT
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0] // Slice the slice to give it zero length. // HL
	printSlice(s)

	s = s[:4] // Extend its length.  // HL
	printSlice(s)

	s = s[2:] // Drop its first two values.  // HL
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// END OMIT
