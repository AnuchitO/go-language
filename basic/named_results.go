package main

import "fmt"

// START OMIT
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// END OMIT
func main() {
	fmt.Println(split(17))
}
