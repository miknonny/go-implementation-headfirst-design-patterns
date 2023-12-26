package main

type Command interface {
	Execute()
}

type LightCommand struct {
	light *Light
}

func (lc *LightCommand) Execute() {
	lc.light.On()
}

func NewLightCommand(l *Light) *LightCommand {
	return &LightCommand{l}
}

type GarageDoorCommand struct {
	garageDoor *GarageDoor
}

func NewGarageDoorCommand(g *GarageDoor) *GarageDoorCommand {
	return &GarageDoorCommand{g}
}

func (g *GarageDoorCommand) Execute() {
	g.garageDoor.Toggle()
}
