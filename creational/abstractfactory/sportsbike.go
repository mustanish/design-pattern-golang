package abstractfactory

// Sportbike to be used globally
type Sportbike struct{}

// GetWheels to be used globally
func (s *Sportbike) GetWheels() int {
	return 2
}

// GetSeats to be used globally
func (s *Sportbike) GetSeats() int {
	return 1
}

// GetType to be used globally
func (s *Sportbike) GetType() int {
	return SportBikeType
}
