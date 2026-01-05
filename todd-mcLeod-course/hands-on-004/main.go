package main

import "fmt"

// go tour step 8 - 10 (https://go.dev/tour/basics/8)

var c, python, java bool
var i2, j2 = 1, 2

func main() {
	// step 8
	var i int
	fmt.Println(i, c, python, java)

	// step 9
	var c2, python2, java2 = true, false, "no!"
	fmt.Println(i2, j2, c2, python2, java2)

	// step 10
	var i3, j3 int = 1, 2
	k := 3
	c3, python3, java3 := true, false, "no!"
	fmt.Println(i3, j3, k, c3, python3, java3)
}
