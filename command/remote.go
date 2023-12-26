package main

type RemoteControl struct {
	slot Command
}

func (r *RemoteControl) PressButton() {
	r.slot.Execute()
}

func (r *RemoteControl) SetCommand(command Command) {
	r.slot = command
}
