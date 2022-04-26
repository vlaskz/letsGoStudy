package adapter

import "fmt"

type Car struct {
	Brand   string
	Color   string
	Started bool
}

func (c *Car) Start() {
	if c.Started {
		fmt.Println("Already Started!")
		return
	}
	fmt.Println("Car is started now!")
	c.Started = true
}

func (c *Car) Accelerate() {
	fmt.Println("Accelerating!")
}

type CarAdapter struct {
	Car *Car
}

func (c *CarAdapter) Move() {
	c.Car.Start()
	c.Car.Accelerate()
}
