package main

import "fmt"

// START OMIT
func main() {
	var a [3]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1], a[2])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// END OMIT
