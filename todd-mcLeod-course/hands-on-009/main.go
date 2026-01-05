package main

import "fmt"

type ByteSize int

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Printf("KB -> %v bytes \t\t\t\t(binary: %b)\n", KB, KB)
	fmt.Printf("MB -> %v bytes \t\t\t(binary: %b)\n", MB, MB)
	fmt.Printf("GB -> %v bytes \t\t\t(binary: %b)\n", GB, GB)
	fmt.Printf("TB -> %v bytes \t\t(binary: %b)\n", TB, TB)
	fmt.Printf("PB -> %v bytes \t(binary: %b)\n", PB, PB)
	fmt.Printf("EB -> %v bytes (binary: %b)\n", EB, EB)
}
