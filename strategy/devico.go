package main

type Device struct {
	Quack Quacker
}

func (d *Device) PerformQuack() {
	d.Quack.Quack()
}
