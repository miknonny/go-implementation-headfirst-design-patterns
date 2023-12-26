package main

import "fmt"

type Light struct {
	on bool
}

func (li *Light) On() {
	li.on = true
	fmt.Println("Lights are turned on.")
}

func (li *Light) Off() {
	li.on = false
	fmt.Println("Lights are turned off")
}

type GarageDoor struct {
	open bool
}

func (g *GarageDoor) Toggle() {
	if g.open {
		g.open = false
		fmt.Println("Garage door closed")
	} else {
		g.open = false
		fmt.Println("Garage door closed")
	}
}
