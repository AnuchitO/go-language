package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// START OMIT
type Path []Point

func (path Path) Distance() float64 { // HL
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i]) // HL
		}
	}
	return sum
} // HL

func main() {
	p := Path{ // HL
		{1, 1}, // HL
		{5, 1}, // HL
		{5, 4}, // HL
		{1, 1}, // HL
	} // HL

	sum := p.Distance() // HL

	fmt.Println("sum of distance", sum)
}

// END OMIT
