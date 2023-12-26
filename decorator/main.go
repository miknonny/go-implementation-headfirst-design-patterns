package main

import "fmt"

func main() {
	order := NewHouseBlend("houseblend", 1.99)

	order = NewMocha("mocha cream", 0.5, order)
	order = NewSoy("soy goodness", 2.0, order)

	fmt.Println(order.Description())
	fmt.Println(order.Cost())
}
