package main

import "fmt"

type rectangle struct {
	width  int
	length int
}

func (rec rectangle) area() int { // HL
	return rec.width * rec.length
} // HL

func main() {
	rec := rectangle{3, 4}
	a := rec.area() // HL
	fmt.Println(a)
}
