package main

import "fmt"

// START OMIT
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 // HL
	fmt.Println(v)
	fmt.Println(p.X)    // HL
	fmt.Println((*p).X) // HL
}

// END OMIT
