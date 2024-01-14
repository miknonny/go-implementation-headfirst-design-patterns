package smarthome

import "fmt"

type Command interface {
	Execute()
}

// Light
type LightOnCommand struct {
	light *Light
}

func (lc *LightOnCommand) Execute() {
	lc.light.On()
}

func NewLightOnCommand(l *Light) *LightOnCommand {
	return &LightOnCommand{l}
}

type LightOffCommand struct {
	light *Light
}

func (lc *LightOffCommand) Execute() {
	lc.light.Off()
}

func NewLightOffCommand(l *Light) *LightOffCommand {
	return &LightOffCommand{l}
}

// Garage Door
type GarageDoorOpenCommand struct {
	garageDoor *GarageDoor
}

func NewGarageDoorOpenCommand(g *GarageDoor) *GarageDoorOpenCommand {
	return &GarageDoorOpenCommand{g}
}

func (g *GarageDoorOpenCommand) Execute() {
	g.garageDoor.Open() // note that we can perform more actions on the recevier of the the garageDoor command.
}

func NewGarageDoorCloseCommand(g *GarageDoor) *GarageDoorCloseCommand {
	return &GarageDoorCloseCommand{g}
}

type GarageDoorCloseCommand struct {
	garageDoor *GarageDoor
}

func (g *GarageDoorCloseCommand) Execute() {
	g.garageDoor.Close()
}

// Stereo command
type StereoOnCommand struct {
	stereo *Stereo
}

func NewSteroOnCommand(s *Stereo) *StereoOnCommand {
	return &StereoOnCommand{s}
}

func (s *StereoOnCommand) Execute() {
	s.stereo.setMode(Cd)
	s.stereo.On()
	s.stereo.setVolume(11)
}

type StereoOffCommand struct {
	stereo *Stereo
}

func NewSteroOffCommand(s *Stereo) *StereoOffCommand {
	return &StereoOffCommand{s}
}

func (s *StereoOffCommand) Execute() {
	s.stereo.Off()
}

// Ceiling Fan
type CeilingFanOnCommand struct {
	ceilingFan *CeilingFan
}

func NewCeilingFanOnCommand(c *CeilingFan) *CeilingFanOnCommand {
	return &CeilingFanOnCommand{
		ceilingFan: c,
	}
}

func (c *CeilingFanOnCommand) Execute() {
	c.ceilingFan.On()
}

type CeilingFanOffCommand struct {
	ceilingFan *CeilingFan
}

func NewCeilingFanOffCommand(c *CeilingFan) *CeilingFanOffCommand {
	return &CeilingFanOffCommand{
		ceilingFan: c,
	}
}

func (c *CeilingFanOffCommand) Execute() {
	c.ceilingFan.Off()
}

type NoCommand struct {
}

func (n *NoCommand) Execute() {
	fmt.Println("program me!")
}
