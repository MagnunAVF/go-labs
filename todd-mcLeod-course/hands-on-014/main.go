package main

import "fmt"

var c int

const X int = 42

func main() {
	var y string = "test"
	c += 1
	z := 10.0

	fmt.Printf("var c -> %d \t\t(type: %T)\n", c, c)
	fmt.Printf("var X -> %d \t(type: %T)\n", X, X)
	fmt.Printf("var y -> %s \t(type: %T)\n", y, y)
	fmt.Printf("var z -> %.2f \t(type: %T)\n", z, z)
}
