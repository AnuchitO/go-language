package main

import "fmt"

func main() {
	defer fmt.Println("world") // HL

	fmt.Println("hello")
}
