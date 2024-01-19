package smarthome

import "fmt"

type Command interface {
	Execute()
	Undo()
}

// Light
type LightOnCommand struct {
	light *Light
}

func (lc *LightOnCommand) Execute() {
	lc.light.On()
}

func (lc *LightOnCommand) Undo() {
	lc.light.Off()
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

func (lc *LightOffCommand) Undo() {
	lc.light.On()
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

func (g *GarageDoorOpenCommand) Undo() {
	g.garageDoor.Close() // note that we can perform more actions on the recevier of the the garageDoor command.
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
func (g *GarageDoorCloseCommand) Undo() {
	g.garageDoor.Open()
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

func (s *StereoOnCommand) Undo() {
	s.stereo.Off()
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

func (s *StereoOffCommand) Undo() {
	s.stereo.On()
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

func (c *CeilingFanOnCommand) Undo() {
	c.ceilingFan.Off()
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

func (c *CeilingFanOffCommand) Undo() {
	c.ceilingFan.On()
}

type NoCommand struct {
}

func (n *NoCommand) Execute() {
	fmt.Println("program me!")
}

func (n *NoCommand) Undo() {

}
