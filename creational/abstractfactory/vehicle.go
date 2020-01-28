package abstractfactory

// Vehicle to be used globally
type Vehicle interface {
	GetWheels() int
	GetSeats() int
}
