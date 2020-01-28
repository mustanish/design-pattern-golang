package abstractfactory

import (
	"errors"
	"fmt"
)

const (
	// SportBikeType to be used globally
	SportBikeType = 1
	// CruiseBikeType to be used globally
	CruiseBikeType = 2
)

// BikeFactory to be used globally
type BikeFactory struct{}

// GetVehicle to be used globally
func (m *BikeFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case SportBikeType:
		return new(Sportbike), nil
	case CruiseBikeType:
		return new(Cruisebike), nil
	default:
		return nil, errors.New(fmt.Sprintf("bike type %d not recognized\n", v))
	}
}
