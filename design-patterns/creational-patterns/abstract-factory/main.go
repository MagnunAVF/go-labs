package main

import "fmt"

func PrintShoeDetails(s IShoe) {
	fmt.Printf("* Shoe - logo: %s , size: %d\n", s.getLogo(), s.getSize())
}

func PrintShirtDetails(s IShoe) {
	fmt.Printf("* Shirt - logo: %s , size: %d\n", s.getLogo(), s.getSize())
}

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	PrintShoeDetails(adidasShoe)
	PrintShirtDetails(adidasShirt)

	PrintShoeDetails(nikeShoe)
	PrintShirtDetails(nikeShirt)
}
