package builder

type Builder interface {
	SetWheels() Builder
	SetSeats() Builder
	SetStructure() Builder
	GetVehicle() VehicleBuilder
}
