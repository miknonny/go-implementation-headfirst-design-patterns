package main

import "githhub.com/miknonny/design_patterns/command/pkg/smarthome"

func main() {
	houseRemote := smarthome.NewRemoteControl()

	// Devices
	livingRoomLight := smarthome.Newlight("Living Room")
	kitchenLight := smarthome.Newlight("kitchen light")
	ceilingFan := smarthome.NewCeilingFan("living Room")
	garageDoor := smarthome.NewGarageDoor("store")

	//Commands
	livingRoomLightOn := smarthome.NewLightOnCommand(livingRoomLight)
	livingRoomLightOff := smarthome.NewLightOffCommand(livingRoomLight)

	kitchenLightOn := smarthome.NewLightOnCommand(kitchenLight)
	kitchenLightOff := smarthome.NewLightOffCommand(kitchenLight)

	ceilingFanOn := smarthome.NewCeilingFanOnCommand(ceilingFan)
	ceilingFanOff := smarthome.NewCeilingFanOffCommand(ceilingFan)

	garageDoorOpen := smarthome.NewGarageDoorOpenCommand(garageDoor)
	garageDoorClose := smarthome.NewGarageDoorCloseCommand(garageDoor)

	// Programming Remote.
	houseRemote.SetCommand(0, livingRoomLightOn, livingRoomLightOff)
	houseRemote.SetCommand(3, garageDoorOpen, garageDoorClose)
	houseRemote.SetCommand(6, ceilingFanOn, ceilingFanOff)
	houseRemote.SetCommand(4, kitchenLightOn, kitchenLightOff)

	houseRemote.OnButtonPushed(0)
	houseRemote.OnButtonPushed(1)
	houseRemote.OnButtonPushed(2)
	houseRemote.OnButtonPushed(3)
	houseRemote.OnButtonPushed(4)
	houseRemote.OnButtonPushed(5)
	houseRemote.OnButtonPushed(6)
}
