package main

import "fmt"

const Pi = 3.14 // HL

func main() {
	const world = "世界" // HL
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")

	const truth = true // HL
	fmt.Println("Go rules?", truth)
}
