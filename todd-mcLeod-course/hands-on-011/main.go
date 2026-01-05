package main

import "fmt"

func main() {
	x := "Test"
	y := 1
	var z float64 = 1.5

	fmt.Printf("x -> %s \t\t(type: %T)\n", x, x)
	fmt.Printf("y -> %d \t\t\t(type: %T)\n", y, y)
	fmt.Printf("z -> %f \t(type: %T)\n", z, z)
}
