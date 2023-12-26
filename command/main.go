package main

import "fmt"

func main() {
	var light Light
	lightCommand := NewLightCommand(&light)

	var garageDoor GarageDoor
	garageDoorCommand := NewGarageDoorCommand(&garageDoor)

	var remote RemoteControl

	remote.SetCommand(lightCommand)
	remote.PressButton()
	remote.PressButton()

	fmt.Println("=========== changed the command on the remote control slot=========")

	remote.SetCommand(garageDoorCommand)
	remote.PressButton()
	remote.PressButton()

}
