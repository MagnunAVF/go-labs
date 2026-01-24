package main

// 3. Liskov Substitution Principle (LSP)
// "Objects of a superclass should be replaceable with objects of its subclasses
// without breaking the application.""
// this translates to: an interface implementation should behave as the caller expects.

// BAD
// A Square is a Rectangle, but it breaks the expectation that you can set Width
// and Height independently.
// type Rectangle struct{ width, height int }

// func (r *Rectangle) SetWidth(w int)  { r.width = w }
// func (r *Rectangle) SetHeight(h int) { r.height = h }

// type Square struct{ Rectangle }

// func (s *Square) SetWidth(w int) { s.width = w; s.height = w } // Breaks behavior

// GOOD
// Don't use embedding to force a relationship that doesn't fit the logic. Treat
// them as separate shapes that fulfill an interface.
type Geometry interface {
	Area() int
}

type Rectangle struct{ W, H int }

func (r Rectangle) Area() int { return r.W * r.H }

type Square struct{ Side int }

func (s Square) Area() int { return s.Side * s.Side }
