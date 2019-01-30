package geometry

import "math"

// START OMIT
type Point struct{ X, Y float64 } // HL

// traditional function
func Distance(p, q Point) float64 { // HL
	return math.Hypot(q.X-p.X, q.Y-p.Y)
} // HL

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 { // HL
	return math.Hypot(q.X-p.X, q.Y-p.Y)
} // HL

// END OMIT
