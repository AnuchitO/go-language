package main

import "fmt"

type rectangle struct {
	width  int
	length int
}

func area(rec rectangle) int {
	return rec.width * rec.length
}

func main() {
	rec := rectangle{3, 4}
	a := area(rec)
	fmt.Println()
}
