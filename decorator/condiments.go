package main

type Mocha struct {
	desscription string
	cost         float64
	beverage     Beverage
}

func NewMocha(description string, cost float64, beverage Beverage) Beverage {
	return &Mocha{description, cost, beverage}
}

func (m *Mocha) Cost() float64 {
	return m.cost + m.beverage.Cost()
}

func (m *Mocha) Description() string {
	return m.desscription + ", " + m.beverage.Description()
}

type Soy struct {
	desscription string
	cost         float64
	beverage     Beverage
}

func NewSoy(description string, cost float64, beverage Beverage) Beverage {
	return &Soy{description, cost, beverage}
}

func (m *Soy) Cost() float64 {
	return m.cost + m.beverage.Cost()
}

func (m *Soy) Description() string {
	return m.desscription + ", " + m.beverage.Description()
}
