package main

import "fmt"

// START OMIT
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},  // HL
	"Google":    {37.42202, -122.08408}, // HL
}

func main() {
	fmt.Println(m)
}

// END OMIT
