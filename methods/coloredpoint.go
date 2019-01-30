package main

import (
	"fmt"
	"math"
)

// START1 OMIT
import "image/color"

type Point struct{ X, Y float64 }
type ColoredPoint struct { // HL
	Point // HL
	Color color.RGBA
} // HL

// END1 OMIT

func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	// START2 OMIT
	var cp ColoredPoint
	cp.X = 1 // HL
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2 // HL
	fmt.Println(cp.Y)
	// END2 OMIT

	// START3 OMIT
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}

	d := p.Distance(q.Point) // HL
	fmt.Println(d)
	p.ScaleBy(2) // HL
	q.ScaleBy(2) // HL

	d = p.Distance(q.Point) // HL
	fmt.Println(d)
	// END3 OMIT
}
