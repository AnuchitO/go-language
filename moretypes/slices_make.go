package main

import "fmt"

// START OMIT
func main() {
	a := make([]int, 5) // HL
	printSlice("a", a)

	b := make([]int, 0, 5) // HL
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// END OMIT
