package main

import (
	"fmt"
	"math"
)

// START OMIT
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // HL
		return v
	} // HL
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

// END OMIT
