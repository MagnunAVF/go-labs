package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// go tour step 11 - 13 (https://go.dev/tour/basics/11)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// step 11
	fmt.Printf("Type: %T \t\t\tValue: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T \t\tValue: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T \tValue: %v\n", z, z)

	// step 12
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// step 13
	var x, y = 3, 4
	var f2 float64 = math.Sqrt(float64(x*x + y*y))
	var z2 uint = uint(f2)
	fmt.Println(x, y, z2)
}
