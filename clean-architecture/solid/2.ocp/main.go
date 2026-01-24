package main

// 2. Open/Closed Principle (OCP)
// "Software entities should be open for extension, but closed for modification.""
// You should be able to add new functionality without changing existing code.

// BAD
// If we want to add a new shape (like a Triangle), we have to modify the
// AreaSum function and add another else if.
// type Rectangle struct{ Width, Height float64 }
// type Circle struct{ Radius float64 }
// func AreaSum(shapes []interface{}) float64 {
// 	var sum float64
// 	for _, shape := range shapes {
// 		if r, ok := shape.(Rectangle); ok {
// 			sum += r.Width * r.Height
// 		} else if c, ok := shape.(Circle); ok {
// 			sum += 3.14 * c.Radius * c.Radius
// 		}
// 	}
// 	return sum
// }

// GOOD
// Use an interface. Now, you can add as many shapes as you want without ever
// touching the AreaSum function.
type Shape interface {
	Area() float64
}

func AreaSum(shapes []Shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.Area()
	}
	return sum
}
