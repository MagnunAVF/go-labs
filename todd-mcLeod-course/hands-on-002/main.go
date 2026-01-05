package main

import (
	"fmt"
	"math"
	"math/rand"
)

// go tour step 1 - 3 (https://go.dev/tour/basics/1)

func main() {
	// step 1
	fmt.Println("My favorite number is", rand.Intn(10))

	// step 2
	fmt.Printf("Now you have %g problems\n", math.Sqrt(7))

	// step 3
	fmt.Println(math.Pi)
}
