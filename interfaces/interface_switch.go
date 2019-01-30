package main

import "fmt"

func main() {
	var i interface{} = "hello"
	switch v := i.(type) {
	case int:
		fmt.Printf("%T %d", v, v)
	case string:
		fmt.Printf("%T %s", v, v)
	default:
		fmt.Println("undefined type")
	}
}
