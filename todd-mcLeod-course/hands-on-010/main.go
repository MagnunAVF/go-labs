package main

import "fmt"

var zero int

func main() {
	t := "test"
	x, y := 1, 2
	var temperature float64 = 1.23
	_, ok := "Something", true

	fmt.Println("zero =", zero)
	fmt.Println("Test string:", t)
	fmt.Printf("x = %d ; y = %d\n", x, y)
	fmt.Printf("Temperature = %f (type: %T)\n", temperature, temperature)
	fmt.Println("ignoring a value ... ok = ", ok)
}
