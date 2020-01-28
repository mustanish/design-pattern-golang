package builder

type BikeBuilder struct {
	vehicle VehicleBuilder
}

func (b *BikeBuilder) SetWheels() Builder {
	b.vehicle.Wheels = 4
	return b
}

func (b *BikeBuilder) SetSeats() Builder {
	b.vehicle.Seats = 5
	return b
}

func (b *BikeBuilder) SetStructure() Builder {
	b.vehicle.Structure = "Bike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleBuilder {
	return b.vehicle
}
