package composition

import "testing"

func TestAthlete_Train(t *testing.T) {
	athlete := AthleteA{}
	athlete.Train()
}

func TestSwimmer_Swim(t *testing.T) {
	localSwim := Swim
	swimmer := SwimmerA{
		Swim: &localSwim,
	}
	swimmer.Athlete.Train()
	(*swimmer.Swim)()
}

func TestAnimal_Swim(t *testing.T) {
	fish := Shark{
		Swim: Swim,
	}
	fish.Eat()
	fish.Swim()
}

func TestSwimmer_Swim2(t *testing.T) {
	swimmer := SwimmerB{
		new(AthleteA),
		new(AthleteB),
	}
	swimmer.Train()
	swimmer.Swim()
}

func TestTree(t *testing.T) {
	root := Tree{
		leafValue: 0,
		right: &Tree{
			leafValue: 5,
			right:     &Tree{6, nil, nil},
		},
		left: &Tree{4, nil, nil},
	}
	if root.right.right.leafValue != 6 {
		t.Errorf("Leaf value must be 6 but it is %d\n", root.right.right.leafValue)
	}
}

func TestSon_GetParentField(t *testing.T) {
	son := Son{}
	GetParentField(son.p)
}
