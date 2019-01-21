package main

import "fmt"

// START OMIT
func main() {
	pow := make([]int, 10)
	for i := range pow { // HL
		pow[i] = 1 << uint(i) // == 2**i
	} // HL
	for _, value := range pow { // HL
		fmt.Printf("%d\n", value)
	} // HL
}

// END OMIT
