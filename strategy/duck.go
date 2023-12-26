package main

import "fmt"

// Every duck has fly and quack.
type Duck struct {
	Fly   Flier
	Quack Quacker
}

func (d *Duck) SetFlyier(f Flier) {
	d.Fly = f
}

func (d *Duck) SetQuaker(f Quacker) {
	d.Quack = f
}

// Part of the duck that never changes. or part of your code that never changes.
func (d *Duck) Swim() {
	fmt.Println("I am swimming")
}

func (d *Duck) Display() {
	fmt.Println("I can display on screen.")
}

type MallardDuck struct {
	Duck
}

func (m *MallardDuck) PerformFly() {
	m.Fly.Fly()
}

func (m *MallardDuck) PerformQuack() {
	m.Quack.Quack()
}
