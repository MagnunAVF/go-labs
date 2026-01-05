package main

import "fmt"

// go tour step 4 - 7 (https://go.dev/tour/basics/4)

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// step 4
	fmt.Println(add(42, 13))

	// step 5
	fmt.Println(add2(42, 13))

	// step 6
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// step 7
	fmt.Println(split(17))
}
