package main

import "fmt"

type Windows struct {
	Printer Printer
}

func (w *Windows) setPrinter(p Printer) {
	w.Printer = p
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.Printer.PrintFile()
}
