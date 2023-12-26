package main

import "fmt"

type Flier interface {
	Fly()
}

type FlyWithWings struct{}

func (f *FlyWithWings) Fly() {
	fmt.Println("I am flying with wings.")
}

type FlyWithNoWings struct{}

func (f *FlyWithNoWings) Fly() {
	fmt.Println("Hey i am flying with no wings.")
}
