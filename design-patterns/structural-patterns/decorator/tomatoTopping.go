package main

type TomatoTopping struct {
	pizza IPizza
}

func (p *TomatoTopping) getPrice() int {
	pizzaPrice := p.pizza.getPrice()
	return pizzaPrice + 7
}
