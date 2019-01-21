package main

import "fmt"

// START OMIT
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Printf("%T => %v\n", primes, primes)

	var s []int = primes[1:4] // HL
	fmt.Printf("%T => %v\n", s, s)
}

// END OMIT
