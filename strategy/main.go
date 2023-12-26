package main

func main() {
	duck := MallardDuck{}
	duck.Fly = &FlyWithNoWings{}
	duck.PerformFly()
	duck.Fly = &FlyWithWings{}
	duck.PerformFly()

	device := Device{}
	device.Quack = &Quack{}
	device.PerformQuack()
}
