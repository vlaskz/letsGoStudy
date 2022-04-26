package adapter

func VehicleFactory(s string) Vehicle {
	switch s {
	case "Car":
		c := Car{}
		return &CarAdapter{&c}
	case "Bike":
		b := Bike{}
		return &BikeAdapter{&b}
	}
	return nil
}
