package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0) // append works on nil slices.
	printSlice(s)

	s = append(s, 1) // The slice grows as needed.
	printSlice(s)

	s = append(s, 2, 3, 4) // We can add more than one element at a time.
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
