package main

import "fmt"

// START OMIT
func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow { // HL
		fmt.Printf("2**%d = %d\n", i, v)
	} // HL
}

// END OMIT
