package smarthome

import (
	"bytes"
	"reflect"
)

const slotSize = 7

type RemoteControl struct {
	onCommands  []Command
	offCommands []Command
	undoCommand Command
}

func NewRemoteControl() *RemoteControl {

	var noCommand NoCommand

	remoteControl := &RemoteControl{
		onCommands:  make([]Command, slotSize),
		offCommands: make([]Command, slotSize),
	}

	// avoid nil pointer error and removes the responsibility of handling null
	// from the client
	for i := 0; i < slotSize; i++ {
		remoteControl.onCommands[i] = &noCommand
		remoteControl.offCommands[i] = &noCommand
	}

	remoteControl.undoCommand = &noCommand

	return remoteControl
}

func (r *RemoteControl) SetCommand(slot int, onCommand Command, offCommand Command) {
	r.onCommands[slot] = onCommand
	r.offCommands[slot] = offCommand
}

func (r *RemoteControl) OnButtonPushed(slot int) {
	r.onCommands[slot].Execute()
	r.undoCommand = r.onCommands[slot]
}

func (r *RemoteControl) OffButtonPushed(slot int) {
	r.offCommands[slot].Execute()
	r.undoCommand = r.offCommands[slot]
}

func (r *RemoteControl) UndoButtonPushed(slot int) {
	r.undoCommand.Undo()
}

func (r *RemoteControl) String() string {
	var buf bytes.Buffer

	buf.WriteString("\n-------- REMOTE CONTROL --------\n")

	// Use the reflect package to get the struct name
	for i := 0; i < len(r.onCommands); i++ {
		onCommandName := reflect.TypeOf(r.onCommands[i]).Name()
		offCommandName := reflect.TypeOf(r.offCommands[i]).Name()
		buf.WriteString(onCommandName + "   " + offCommandName + "\n")
	}

	return buf.String()
}
