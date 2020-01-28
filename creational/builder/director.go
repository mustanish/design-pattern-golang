package builder

// Director is exported to be used globally
type Director struct {
	builder Builder
}

// Construct is exported to be used globally
func (d *Director) Construct() {
	d.builder.SetSeats().SetWheels().SetStructure()
}

// SetBuilder is exported to be used globally
func (d *Director) SetBuilder(b Builder) {
	d.builder = b
}
