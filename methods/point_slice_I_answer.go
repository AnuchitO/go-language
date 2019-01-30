package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) { // HL
	p.X *= factor
	p.Y *= factor
} // HL

// START OMIT
func main() {
	r := &Point{1, 2} // HL
	r.ScaleBy(2)      // HL
	fmt.Println(*r)   // "{2, 4}"

	// or this:
	p := Point{1, 2} // HL
	pptr := &p       // HL
	pptr.ScaleBy(2)
	fmt.Println(p) // "{2, 4}"

	// or this:
	p = Point{1, 2} // HL
	(&p).ScaleBy(2) // HL
	fmt.Println(p)  // "{2, 4}"
}

// END OMIT
