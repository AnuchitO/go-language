package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	t := time.Now()
	switch { // HL
	case t.Hour() < 12: // HL
		fmt.Println("Good morning!")
	case t.Hour() < 17: // HL
		fmt.Println("Good afternoon.")
	default: // HL
		fmt.Println("Good evening.")
	} // HL
}

// END OMIT
