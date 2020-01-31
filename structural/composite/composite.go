package composition

type trainer interface {
	Train()
}

// SwimmerA exported to be used globally
type SwimmerA struct {
	Athlete AthleteA
	Swim    *func()
}

// AthleteA exported to be used globally
type AthleteA struct{}

// Train exported to be used globally
func (a *AthleteA) Train() {
	println("Training")
}

// Swim exported to be used globally
func Swim() {
	println("Swimming Function!")
}

// Animal exported to be used globally
type Animal struct{}

// Eat exported to be used globally
func (a *Animal) Eat() {
	println("Eating")
}

// Shark exported to be used globally
type Shark struct {
	Animal
	Swim func()
}

type swimmer interface {
	Swim()
}

// AthleteB exported to be used globally
type AthleteB struct{}

// Swim exported to be used globally
func (a *AthleteB) Swim() {
	println("Swimming Method!")
}

// SwimmerB exported to be used globally
type SwimmerB struct {
	trainer
	swimmer
}

// Tree exported to be used globally
type Tree struct {
	leafValue int
	right     *Tree
	left      *Tree
}

type parent struct {
	someField int
}

// Son exported to be used globally
type Son struct {
	p parent
}

// GetParentField exported to be used globally
func GetParentField(p parent) int {
	return p.someField
}
