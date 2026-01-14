package main

import "fmt"

func main() {
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}

	macComputer.setPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.setPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	windowsComputer := &Windows{}

	windowsComputer.setPrinter(hpPrinter)
	windowsComputer.Print()
	fmt.Println()

	windowsComputer.setPrinter(epsonPrinter)
	windowsComputer.Print()
	fmt.Print()
}
