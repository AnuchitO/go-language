package main

import (
	"fmt"
	"math"
)

// START OMIT
func sqrt(x float64) string {
	if x < 0 { // HL
		return sqrt(-x) + "i"
	} // HL
	return fmt.Sprint(math.Sqrt(x))
}

// END OMIT

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
