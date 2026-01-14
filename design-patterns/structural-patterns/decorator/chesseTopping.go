package main

type CheeseTopping struct {
	pizza IPizza
}

func (p *CheeseTopping) getPrice() int {
	pizzaPrice := p.pizza.getPrice()
	return pizzaPrice + 10
}
