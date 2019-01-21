package main

import "fmt"

// START OMIT
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil { // HL
		fmt.Println("nil!")
	} // HL
}

// END OMIT
