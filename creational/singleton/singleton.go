package singleton

// Singleton exported to be used globally
type Singleton struct {
	count int
}

var instance *Singleton

// GetInstance exported to be used globally
func GetInstance() *Singleton {
	if instance == nil {
		instance = new(Singleton)
	}
	return instance
}

// AddOne exported to be used globally
func (s *Singleton) AddOne() int {
	s.count++
	return s.count
}

// GetCount exported to be used globally
func (s *Singleton) GetCount() int {
	return s.count
}
