package abstractfactory

import (
	"errors"
	"fmt"
)

// VehicleFactory to be used globally
type VehicleFactory interface {
	GetVehicle(v int) (Vehicle, error)
}

const (
	// CarFactoryType to be used globally
	CarFactoryType = 1
	// BikeFactoryType to be used globally
	BikeFactoryType = 2
)

// GetVehicleFactory to be used globally
func GetVehicleFactory(factoryType int) (VehicleFactory, error) {
	switch factoryType {
	case CarFactoryType:
		return new(CarFactory), nil
	case BikeFactoryType:
		return new(BikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("factory with id %d not recognized", factoryType))
	}
}
