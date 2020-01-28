package abstractfactory

// Cruisebike to be used globally
type Cruisebike struct{}

// GetWheels to be used globally
func (c *Cruisebike) GetWheels() int {
	return 2
}

// GetSeats to be used globally
func (c *Cruisebike) GetSeats() int {
	return 2
}

// GetType to be used globally
func (c *Cruisebike) GetType() int {
	return CruiseBikeType
}
