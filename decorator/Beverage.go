package main

// Here are the base Classes to be decorated.
type Beverage interface {
	Cost() float64
	Description() string
}

type HouseBlend struct {
	description string
	cost        float64
}

func (h *HouseBlend) Cost() float64 {
	return h.cost
}

func (h *HouseBlend) Description() string {
	return h.description
}

func NewHouseBlend(description string, cost float64) Beverage {
	return &HouseBlend{description, cost}
}

type Expresso struct {
	description string
	cost        float64
}

func (e *Expresso) Cost() float64 {
	return e.cost
}

func (e *Expresso) Description() string {
	return e.description
}

func NewExpresso(description string, cost float64) Beverage {
	return &Expresso{description, cost}
}
