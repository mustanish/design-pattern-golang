package abstractfactory

import "testing"

func TestBikeFactory(t *testing.T) {
	bikeFactory, err := GetVehicleFactory(BikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	bikeType, err := bikeFactory.GetVehicle(SportBikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Bike vehicle has %d wheels and %d seats\n", bikeType.GetWheels(), bikeType.GetSeats())

	sportBike, ok := bikeType.(Bike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Sport motorbike has type %d\n", sportBike.GetType())

	bikeType, err = bikeFactory.GetVehicle(CruiseBikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Bike vehicle has %d wheels and %d seats\n", bikeType.GetWheels(), bikeType.GetSeats())

	cruiseBike, ok := bikeType.(Bike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Cruise motorbike has type %d\n", cruiseBike.GetType())

	bikeType, err = bikeFactory.GetVehicle(3)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCarFactory(t *testing.T) {
	carFactory, err := GetVehicleFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}
	carType, err := carFactory.GetVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Car vehicle has %d seats and %d wheels\n", carType.GetSeats(), carType.GetWheels())

	luxuryCar, ok := carType.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Luxury car has %d doors.\n", luxuryCar.GetDoors())

	carType, err = carFactory.GetVehicle(FamiliarCarType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Car vehicle has %d seats and %d wheels\n", carType.GetSeats(), carType.GetWheels())

	familyCar, ok := carType.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Familiy car has %d doors.\n", familyCar.GetDoors())

	carType, err = carFactory.GetVehicle(3)
	if err != nil {
		t.Fatal(err)
	}
}
