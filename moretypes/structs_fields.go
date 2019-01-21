package main

import "fmt"

type Vertex struct {
	X int // HL
	Y int // HL
}

func main() {
	v := Vertex{1, 2}
	v.X = 4          // HL
	fmt.Println(v.X) // HL
}
