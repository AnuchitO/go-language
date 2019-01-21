package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday { // HL
	case today + 0: // HL
		fmt.Println("Today.")
	case today + 1: // HL
		fmt.Println("Tomorrow.")
	default: // HL
		fmt.Println("Too far away.")
	}
}

// END OMIT
