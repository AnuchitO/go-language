package main

import (
	"fmt"
	"math"
)

import "image/color"

type Point struct{ X, Y float64 }
type ColoredPoint struct { // HL
	Point // HL
	Color color.RGBA
} // HL

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

	methodValue()
	methodExpression()
}

func methodValue() {
	// START_VALUE OMIT
	p := Point{1, 2}
	q := Point{4, 6}
	distanceFromP := p.Distance // HL
	fmt.Println(distanceFromP(q))

	var origin Point
	fmt.Println(distanceFromP(origin))
	scaleP := p.ScaleBy // HL
	scaleP(2)
	scaleP(3)
	scaleP(10)
	// END_VALUE OMIT
}

func methodExpression() {
	// START_EXP OMIT
	p := Point{1, 2}
	q := Point{4, 6}
	distance := Point.Distance // HL
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance) // "func(Point, Point) float64"

	scale := (*Point).ScaleBy // HL
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale) // "func(*Point, float64)"
	// END_EXP OMIT
}
