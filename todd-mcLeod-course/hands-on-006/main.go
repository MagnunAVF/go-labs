package main

import (
	"fmt"
)

// go tour step 14 - 15 (https://go.dev/tour/basics/14)

const Pi = 3.14

func main() {
	// step 14
	v := 42.0
	fmt.Printf("v is of type %T\n", v)

	// step 15
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
