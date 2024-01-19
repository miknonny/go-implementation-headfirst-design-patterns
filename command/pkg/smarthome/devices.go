package smarthome

import (
	"fmt"
)

type StereoMode int

const (
	Cd StereoMode = iota
	Dvd
	Radio
)

type Light struct {
	location string
	on       bool
}

func Newlight(location string) *Light {
	return &Light{location: location}
}

func (li *Light) On() {
	li.on = true
	fmt.Printf("%s Lights was turned on\n", li.location)
}

func (li *Light) Off() {
	li.on = false
	fmt.Printf("%s Lights was turned off\n", li.location)
}

type GarageDoor struct {
	location string
	open     bool
}

func NewGarageDoor(location string) *GarageDoor {
	return &GarageDoor{}
}

func (g *GarageDoor) Open() {
	g.open = true
	fmt.Printf("%s GarageDoor was opened\n", g.location)
}

func (g *GarageDoor) Close() {
	g.open = false
	fmt.Printf("%s GarageDoor was closed", g.location)
}

type Stereo struct {
	location string
	on       bool
	mode     StereoMode
	volume   int
}

func NewStereo(location string) *Stereo {
	return &Stereo{location: location}
}

func (s *Stereo) On() {
	s.on = true
	s.setVolume(6)
	fmt.Println("turned on stereo")
}

func (s *Stereo) Off() {
	s.on = false
	fmt.Printf("%v Stereo was turned on", s.location)
}

func (s *Stereo) setMode(mode StereoMode) {
	s.mode = mode
	fmt.Printf("%v Stereo was turned off", s.mode)
}

func (s *Stereo) setVolume(vol int) {
	s.volume = vol
	fmt.Println("")
}

type CeilingFan struct {
	location string
	on       bool
}

func NewCeilingFan(location string) *CeilingFan {
	return &CeilingFan{location: location}
}
func (c *CeilingFan) On() {
	c.on = true
	fmt.Printf("%s ceiling fan was turned on", c.location)
}

func (c *CeilingFan) Off() {
	c.on = false
	fmt.Printf("%s ceiling fan was turned off", c.location)
}
