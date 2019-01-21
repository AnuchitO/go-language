package main

import "fmt"

func adder() (func() int, func() int) {
	sum := 0
	return func() int {
			sum += 1
			return sum
		},
		func() int {
			return sum
		}
}

func main() {
	inc, cur := adder()
	fmt.Println(cur())

	fmt.Println(inc())
	fmt.Println(inc())

	fmt.Println(cur())
}
