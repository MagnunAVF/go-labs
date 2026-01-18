package main

import "math"

// The Bad: The "Hybrid" Struct This struct tries to be a simple data container
// but also hides business logic, leading to confusion about its purpose.
// type Rectangle struct {
// 	Width  float64
// 	Height float64
// }
// // Logic is tied to the data structure
// func (r *Rectangle) Area() float64 {
// 	return r.Width * r.Height
// }

// The Good: Pure Interfaces (Polymorphism) We hide the "how" behind an
// interface. The calling code doesn't care if it's a Square or a Circle; it
// just knows it can calculate an Area().
type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
