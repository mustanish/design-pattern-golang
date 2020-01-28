package abstractfactory

import (
	"errors"
	"fmt"
	"log"
)

const (
	// LuxuryCarType to be used globally
	LuxuryCarType = 1
	// FamiliarCarType to be used globally
	FamiliarCarType = 2
)

// CarFactory to be used globally
type CarFactory struct{}

// GetVehicle to be used globally
func (f *CarFactory) GetVehicle(v int) (Vehicle, error) {
	log.Println(v)
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamiliarCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("car type %d not recognized\n", v))
	}
}
