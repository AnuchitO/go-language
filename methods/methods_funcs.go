package main

import (
	"fmt"
	"math"
)

// START OMIT
type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 { // HL
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
} // HL

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v)) // HL
}

// END OMIT
