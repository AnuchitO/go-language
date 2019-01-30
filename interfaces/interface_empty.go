package main

import "fmt"

func main() {
	var v interface{}
	v = 1
	fmt.Printf("%T %#v\n", v, v)
	v = "1"
	fmt.Printf("%T %#v\n", v, v)
}
