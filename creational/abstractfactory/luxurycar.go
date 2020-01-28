package abstractfactory

// LuxuryCar to be used globally
type LuxuryCar struct{}

// GetDoors to be used globally
func (l *LuxuryCar) GetDoors() int {
	return 4
}

// GetWheels to be used globally
func (l *LuxuryCar) GetWheels() int {
	return 4
}

// GetSeats to be used globally
func (l *LuxuryCar) GetSeats() int {
	return 5
}
