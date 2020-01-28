package builder

type BusBuilder struct {
	vehicle VehicleBuilder
}

func (b *BusBuilder) SetWheels() Builder {
	b.vehicle.Wheels = 4
	return b
}

func (b *BusBuilder) SetSeats() Builder {
	b.vehicle.Seats = 6
	return b
}

func (b *BusBuilder) SetStructure() Builder {
	b.vehicle.Structure = "Bus"
	return b
}

func (b *BusBuilder) GetVehicle() VehicleBuilder {
	return b.vehicle
}
