package main

import "fmt"

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	fmt.Printf("* Gun: %s , Power: %d\n", ak47.getName(), ak47.getPower())
	fmt.Printf("* Gun: %s , Power: %d\n", musket.getName(), musket.getPower())

	_, err := getGun("wrong type")
	if err != nil {
		fmt.Printf("Error in weapon creation: %s\n", err)
	}
}
