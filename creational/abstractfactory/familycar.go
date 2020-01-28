package abstractfactory

// FamilyCar to be used globally
type FamilyCar struct{}

// GetDoors to be used globally
func (l *FamilyCar) GetDoors() int {
	return 5
}

// GetWheels to be used globally
func (l *FamilyCar) GetWheels() int {
	return 4
}

// GetSeats to be used globally
func (l *FamilyCar) GetSeats() int {
	return 5
}
