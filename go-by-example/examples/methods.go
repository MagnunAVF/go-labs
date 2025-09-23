package examples

import "fmt"

type rectA struct {
	width, height int
}

func (r *rectA) area() int {
	return r.width * r.height
}

func (r rectA) perim() int {
	return 2*r.width + 2*r.height
}

func Methods() {
	r := rectA{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
