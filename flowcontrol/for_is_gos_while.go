package main

import "fmt"

// START OMIT
func main() {
	sum := 1
	for sum < 1000 { // HL
		sum += sum
	} // HL
	fmt.Println(sum)
}

// END OMIT
