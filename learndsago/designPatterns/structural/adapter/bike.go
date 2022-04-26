package adapter

import "fmt"

type Bike struct {
	Brand string
	Color string
}

func (b *Bike) Run() {
	fmt.Println("Bike is Running!")
}

type BikeAdapter struct {
	Bike *Bike
}

func (b *BikeAdapter) Move() {
	b.Bike.Run()
}
