package main

import "fmt"

type iSportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}
	if brand == "nike" {
		return &Nike{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}
