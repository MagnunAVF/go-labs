package main

// 4. Interface Segregation Principle (ISP)
// "Clients should not be forced to depend upon interfaces that they do not use."

// Go favors small, focused interfaces. "The bigger the interface, the weaker the
// abstraction." — Rob Pike.

// BAD
// A SimplePrinter is forced to implement Scan and Fax even if it can't do those things.
// type MultiFunctionDevice interface {
// 	Print()
// 	Scan()
// 	Fax()
// }

// type SimplePrinter struct{}

// func (s SimplePrinter) Print() {}
// func (s SimplePrinter) Scan()  { panic("not supported") } // Bad!

// GOOD
// Split the interface into smaller pieces.
type Printer interface{ Print() }
type Scanner interface{ Scan() }

type SimplePrinter struct{}

func (s SimplePrinter) Print() { /* Logic */ } // Only implements what it needs
