package main

import "fmt"

type Observer interface {
	Update(float64, float64, float64)
	Display()
}

type CurrentCondition struct {
	ID                              int
	subject                         Subject
	temperature, humidity, pressure float64
}

func (c *CurrentCondition) Update(temp, humidity, pressure float64) {
	c.temperature, c.humidity, c.pressure = temp, humidity, pressure
	// when  the values changed we
	c.Display()
}

// We want to be able to unregister ourselves that is why we stored a reference to the subject.
func NewCurrentCondition(id int, s Subject) Observer {
	return &CurrentCondition{
		ID:      id,
		subject: s,
	}
}

func (c *CurrentCondition) UnRegister() {
	c.subject.RemoveObserver(c)
}

func (c *CurrentCondition) Display() {
	fmt.Printf("%.2f : %.2f : %.2f\n", c.temperature, c.humidity, c.pressure)
}
