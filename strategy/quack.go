package main

import "fmt"

type Quacker interface {
	Quack()
}

type Quack struct{}

func (q *Quack) Quack() {
	fmt.Println("I can quack")
}

type Squeak struct{}

func (s *Squeak) Quack() {
	fmt.Println("I am squeaking")
}

type MuteQuack struct{}

func (m *MuteQuack) Quack() {
	fmt.Println("I lost my voice")
}
