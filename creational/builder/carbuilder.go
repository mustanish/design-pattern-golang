package builder

// CarBuilder is exported to be used globally
type CarBuilder struct {
	vehicle VehicleBuilder
}

// SetWheels is exported to be used globally
func (c *CarBuilder) SetWheels() Builder {
	c.vehicle.Wheels = 4
	return c
}

// SetSeats is exported to be used globally
func (c *CarBuilder) SetSeats() Builder {
	c.vehicle.Seats = 7
	return c
}

// SetStructure is exported to be used globally
func (c *CarBuilder) SetStructure() Builder {
	c.vehicle.Structure = "Car"
	return c
}

// GetVehicle is exported to be used globally
func (c *CarBuilder) GetVehicle() VehicleBuilder {
	return c.vehicle
}
