package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 { // HL
	return fn(3, 4)
} // HL

func main() {
	hypot := func(x, y float64) float64 { // HL
		return math.Sqrt(x*x + y*y)
	} // HL
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))    // HL
	fmt.Println(compute(math.Pow)) // HL
}
